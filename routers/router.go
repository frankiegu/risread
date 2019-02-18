package routers

import (
	"risread/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
	),
)
	beego.AddNamespace(ns)
}
