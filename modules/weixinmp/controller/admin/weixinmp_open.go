package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gcslaoli/cool-admin-go-modules/modules/weixinmp/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type WeixinmpOpenController struct {
	*cool.ControllerSimple
}

func init() {
	var weixinmp_open_controller = &WeixinmpOpenController{
		&cool.ControllerSimple{
			Perfix: "/admin/weixinmp/open",
		},
	}
	// 注册路由
	cool.RegisterControllerSimple(weixinmp_open_controller)
}

// MessageGetReq 消息请求
type MessageGetReq struct {
	g.Meta    `path:"/message" method:"GET"`
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Echostr   string `json:"echostr"`
}

// MessageGetRes 消息响应
type MessageGetRes struct {
}

// MessageGet 消息接收 用于微信公众号接入认证
func (c *WeixinmpOpenController) MessageGet(ctx context.Context, req *MessageGetReq) (res *MessageGetRes, err error) {
	r := g.RequestFromCtx(ctx)
	WeixinmpService := service.NewWeixinmpService()
	// 验证签名
	if !WeixinmpService.CheckSignature(ctx, req.Signature, req.Timestamp, req.Nonce) {
		return nil, gerror.New("签名错误")
	}
	r.Response.WriteExit(req.Echostr)
	return
}

// MessagePostReq 消息请求
type MessagePostReq struct {
	g.Meta `path:"/message" method:"POST"`
}

// MessagePostRes 消息响应
type MessagePostRes struct {
}

// MessagePost 消息接收 用于接收微信公众号消息
func (c *WeixinmpOpenController) MessagePost(ctx context.Context, req *MessagePostReq) (res *MessagePostRes, err error) {
	r := g.RequestFromCtx(ctx)
	reqmap := r.GetMap()
	g.Dump(reqmap)

	r.Response.WriteExit("success")
	return
}
