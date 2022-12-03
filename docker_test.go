package main_test

import (
	"testing"

	"github.com/gcslaoli/cool-admin-go-modules/modules/docker/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestDocker(t *testing.T) {
	t.Log("TestDocker")
	ctx := gctx.New()
	docker := service.NewDockerService()
	info, err := docker.Client.Info(ctx)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(info)
		g.Dump(info)
	}

}
