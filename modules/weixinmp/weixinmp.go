package demo

import (
	_ "github.com/gcslaoli/cool-admin-go-modules/modules/weixinmp/controller"
	_ "github.com/gcslaoli/cool-admin-go-modules/modules/weixinmp/middleware"
	"github.com/gcslaoli/cool-admin-go-modules/modules/weixinmp/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	WeiXinMpService = service.NewWeixinmpService()
)

func init() {
	ctx := gctx.GetInitCtx()
	g.Log().Debug(ctx, "modules weixinmp init start ...")
	g.Log().Debug(ctx, "modules weixinmp init finished ...")
}
