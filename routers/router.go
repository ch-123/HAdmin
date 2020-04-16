package routers

import (
	"admin-api/controllers/admin"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {

	beego.Include(&admin.Admin{})

	beego.Any("/*", func(Ctx *context.Context) {
		Ctx.WriteString("404")
	})

}
