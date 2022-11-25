package main

import (
	_ "github.com/gcslaoli/cool-admin-go-modules/internal/packed"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"

	_ "github.com/gcslaoli/cool-admin-go-modules/modules"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gcslaoli/cool-admin-go-modules/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
