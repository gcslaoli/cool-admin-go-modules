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

	res, err := cool.CacheManager.GetOrSetFunc(ctx, "weixinmp:accesstoken", getFunc, 3600*time.Second)
	if err != nil {
		return "", err
	}
	accessToken = res.String()
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

// UserGet 获取用户列表
func (s *WeixinmpService) UserGet(ctx g.Ctx, nextOpenid string) (total int, count int, openid []string, next_openid string, err error) {
	accessToken, err := s.GetAccessToken(ctx)
	if err != nil {
		return 0, 0, nil, "", err
	}
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + accessToken + "&next_openid=" + nextOpenid
	type Result struct {
		Total int `json:"total"`
		Count int `json:"count"`
		Data  struct {
			Openid []string `json:"openid"`
		} `json:"data"`
		NextOpenid string `json:"next_openid"`
		Errcode    int    `json:"errcode"`
		Errmsg     string `json:"errmsg"`
	}
	var result Result
	g.Client().GetVar(ctx, url).Scan(&result)
	if result.Errcode != 0 {
		err = gerror.New(result.Errmsg)
		return
	}
	total = result.Total
	count = result.Count
	openid = result.Data.Openid
	next_openid = result.NextOpenid
	return
}

// UserInfo 用户信息
type UserInfo struct {
	Subscribe      int    `json:"subscribe"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	Openid         string `json:"openid"`          // 用户的标识，对当前公众号唯一
	Language       string `json:"language"`        // 用户的语言，简体中文为zh_CN
	SubscribeTime  int    `json:"subscribe_time"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid        string `json:"unionid"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string `json:"remark"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注）
	Groupid        int    `json:"groupid"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []int  `json:"tagid_list"`      // 用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene"` // 返回用户关注的渠道来源
	QrScene        int    `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）
}

// GetUserInfo 获取用户信息
func (s *WeixinmpService) GetUserInfo(ctx g.Ctx, openid string, lang ...string) (userInfo *UserInfo, err error) {
	if openid == "" {
		err = gerror.New("openid不能为空")
		return
	}
	if len(lang) == 0 {
		lang = append(lang, "zh_CN")
	}
	accessToken, err := s.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/cgi-bin/user/info?access_token=" + accessToken + "&openid=" + openid + "&lang=" + lang[0]
	type Result struct {
		*UserInfo
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
	var result Result
	g.Client().GetVar(ctx, url).Scan(&result)
	if result.Errcode != 0 {
		err = gerror.New(result.Errmsg)
		return
	}
	userInfo = result.UserInfo
	return
}

//

// BatchGetUserInfo 批量获取用户信息
func (s *WeixinmpService) BatchGetUserInfo(ctx g.Ctx, user_list []map[string]string) (user_info_list []*UserInfo, err error) {
	accessToken, err := s.GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	url := "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=" + accessToken
	type Result struct {
		UserInfoList []*UserInfo `json:"user_info_list"`
		Errcode      int         `json:"errcode"`
		Errmsg       string      `json:"errmsg"`
	}
	var result Result
	g.Client().ContentJson().PostVar(ctx, url, g.Map{"user_list": user_list}).Scan(&result)
	if result.Errcode != 0 {
		err = gerror.New(result.Errmsg)
		return
	}
	user_info_list = result.UserInfoList
	return
}

// UpdateRemark 设置用户备注名
func (s *WeixinmpService) UpdateRemark(ctx g.Ctx, openid, remark string) (err error) {
	accessToken, err := s.GetAccessToken(ctx)
	if err != nil {
		return err
	}
	url := "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=" + accessToken
	type Result struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
	var result Result
	g.Client().ContentJson().PostVar(ctx, url, g.Map{"openid": openid, "remark": remark}).Scan(&result)
	if result.Errcode != 0 {
		err = gerror.New(result.Errmsg)
		return
	}
	return
}
