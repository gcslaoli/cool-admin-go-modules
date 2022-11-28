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
	total, count, openid, next_openid, err := WeiXinMpService.UserGet(ctx, "")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("total: %d, count: %d, openid: %s, next_openid: %s", total, count, openid, next_openid)
	}
	userInfo, err := WeiXinMpService.GetUserInfo(ctx, openid[0], "zh_CN")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("userInfo: %+v", userInfo)
	}
	userInfoList, err := WeiXinMpService.BatchGetUserInfo(ctx, []map[string]string{{"openid": openid[0], "lang": "zh_CN"}})
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("userInfoList: %+v", userInfoList[0])
	}
	// 设置用户备注名
	err = WeiXinMpService.UpdateRemark(ctx, openid[0], "test")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("UpdateRemark success")
	}

}
