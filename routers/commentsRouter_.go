package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["admin-api/controllers/admin:Admin"] = append(beego.GlobalControllerRouter["admin-api/controllers/admin:Admin"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/admin/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
