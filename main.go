package main

import (
	_ "risread/routers"

	"github.com/astaxie/beego"
)

func init() {

}

func main() {

	configApp()

	beego.Run()
}

// config this app
func configApp() {
	beego.BConfig.AppName = "risRead"
	beego.BConfig.ServerName = "risRead"

	// swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}
