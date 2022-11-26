package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// 微信公众号配置
type weinxinMpConfig struct {
	AppId     string `json:"app_id,omitempty"`
	AppSecret string `json:"app_secret,omitempty"`
	Token     string `json:"token,omitempty"`
}

// 微信公众号配置
var WeixinMp weinxinMpConfig

func init() {
	ctx := gctx.GetInitCtx()
	WeixinMp = weinxinMpConfig{
		AppId:     g.Cfg().MustGetWithEnv(ctx, "modules.weixinmp.appid").String(),
		AppSecret: g.Cfg().MustGetWithEnv(ctx, "modules.weixinmp.appsecret").String(),
		Token:     g.Cfg().MustGetWithEnv(ctx, "modules.weixinmp.token").String(),
	}
}
