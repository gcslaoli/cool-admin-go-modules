package service

import (
	"time"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gcslaoli/cool-admin-go-modules/modules/weixinmp/config"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type WeixinmpService struct {
}

func NewWeixinmpService() *WeixinmpService {
	return &WeixinmpService{}
}

// GetAccessToken 获取微信公众号access_token 缓存1小时
func (s *WeixinmpService) GetAccessToken(ctx g.Ctx) (accessToken string, err error) {
	// cool.CacheManager.GetOrSetFuncLock(ctx, "weixinmp_access_token", func() (interface{}, error) {
	//  getFunc 获取access_token的方法
	getFunc := func(ctx g.Ctx) (value interface{}, err error) {
		// 获取access_token
		url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + config.WeixinMp.AppId + "&secret=" + config.WeixinMp.AppSecret
		type Result struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
			Errcode     int    `json:"errcode"`
			Errmsg      string `json:"errmsg"`
		}
		var result Result
		g.Client().GetVar(ctx, url).Scan(&result)
		if result.Errcode != 0 {
			return nil, gerror.New(result.Errmsg)
		}
		if result.AccessToken == "" {
			return nil, gerror.New("获取access_token失败")
		}
		value = result.AccessToken
		return
	}

	result, err := cool.CacheManager.GetOrSetFunc(ctx, "weixinmp:accesstoken", getFunc, 3600*time.Second)
	if err != nil {
		return "", err
	}
	accessToken = result.String()
	return
}

// GetApiDomainIP 获取微信公众号服务器IP地址
func (s *WeixinmpService) GetApiDomainIP(ctx g.Ctx) (ipList []string, err error) {
	accessToken, err := s.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=" + accessToken
	type Result struct {
		IpList  []string `json:"ip_list"`
		Errcode int      `json:"errcode"`
		Errmsg  string   `json:"errmsg"`
	}
	var result Result
	g.Client().GetVar(ctx, url).Scan(&result)
	if result.Errcode != 0 {
		return nil, gerror.New(result.Errmsg)
	}
	ipList = result.IpList

	return
}

// CheckSignature 微信公众号消息校验
func (s *WeixinmpService) CheckSignature(ctx g.Ctx, signature string, timestamp string, nonce string) (isWX bool) {
	// 获取token
	token := config.WeixinMp.Token
	// 将三个参数并入一个数组
	var arr = garray.NewStrArray()
	arr.SetArray([]string{token, timestamp, nonce})
	// 对数组进行排序
	arr.Sort()
	// 将排序后的数组拼接成字符串
	str := arr.Join("")
	// 对字符串进行sha1加密
	sha1Str := gsha1.Encrypt(str)
	// 将加密后的字符串与signature进行对比
	if sha1Str == signature {
		isWX = true
	}
	return
}
