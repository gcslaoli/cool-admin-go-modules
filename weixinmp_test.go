package main_test

import (
	"testing"

	"github.com/gcslaoli/cool-admin-go-modules/modules/weixinmp/service"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestWeixinmp(t *testing.T) {
	t.Log("TestWeixinmp")
	var ctx = gctx.New()
	WeiXinMpService := service.NewWeixinmpService()
	token, err := WeiXinMpService.GetAccessToken(ctx)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("token: %s", token)
	}
	ipList, err := WeiXinMpService.GetApiDomainIP(ctx)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("ipList: %s", ipList)
	}

}
