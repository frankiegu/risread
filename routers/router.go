package routers

import (
	"risread/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			// 登录
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			// 注册
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
		),
	)
	beego.AddNamespace(ns)
}
