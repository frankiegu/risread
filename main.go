package main

import (
	_ "risread/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"os"
	_"github.com/go-sql-driver/mysql"

	"risread/models"
)

func init() {
	initDataBaseMysql(true,"mysql",&models.User{})
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

// 初始化mysql 数据库
// 如果遇到错误,打印错误日志信息,退出
func initDataBaseMysql(debug bool ,driveName string ,models ...interface{}) {

	orm.Debug = debug
	var err error
	err = orm.RegisterDriver(driveName, orm.DRMySQL)
	// 如果注册数据库启动失败,打印错误日志,退出程序
	if err != nil {
		fmt.Println("error orm.RegisterDriver ", err)
		os.Exit(2)
	}
	// 获取数据库的地址.. 这里的地址从配置文件中读取
	dataSource := beego.AppConfig.String("source")
	// 注册 MySQL数据库, 使用的default别名
	err = orm.RegisterDataBase("default", "mysql", dataSource, 30)
	if err != nil {
		// 注册数据库失败, 打印错误信息退出程序
		fmt.Println("error orm.RegisterDataBase", err)
		os.Exit(2)
	}

	// 注册数据库的模型
	orm.RegisterModel(models...)
}
