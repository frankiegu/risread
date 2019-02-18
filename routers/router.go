package routers

import (
	"risread/controllers"

	"github.com/astaxie/beego"
	"log"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			// 登录
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			// 注册
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
			// 上传文件书籍
			beego.NSRouter("/cstorage",&controllers.UserController{},"post:CloudStorage"),
			// 查看所有上传的书籍
			beego.NSRouter("/viewbooks",&controllers.UserController{},"get:PullCloudStorage"),
		),

	)
	beego.AddNamespace(ns)

	beego.InsertFilter("/static/upload/*",beego.BeforeStatic, func(context *context.Context) {
		log.Printf("before router filter ")
		auth:=context.Input.Header(controllers.JwtAuth)

		log.Println("auth:",auth)
		// todo
	})


}
