package main

import (
	_ "risread/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"

	"risread/models"
	"time"
)

func init() {
	initDataBaseMysql(true, "mysql",
		&models.User{}, &models.UploadBook{}, &models.BookListType{},
		&models.BookList{}, &models.BookInfo{}, &models.BookType{},
		&models.BookAttainment{}, &models.BookListFavorite{}, &models.BookListComment{},
		&models.BookInfoComment{},
	)
	//testInsertBookList()
	//testInsertBookInfo()
	//testReadBookInfos()
	//testReadAllBookListTypes()
}

func main() {

	configApp()
	// 设置静态文件
	beego.SetStaticPath("/static/*", "static")
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
func initDataBaseMysql(debug bool, driveName string, models ...interface{}) {

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

	// 测试多对多插入数据
	//testInsertM2mBookListBookInfo()
	//// 数据库别名
	//name := "default"
	//
	//// drop table 后再建表
	//force := true
	//
	//// 打印执行过程
	//verbose := true
	//
	//// 遇到错误立即返回
	//err = orm.RunSyncdb(name, force, verbose)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//orm.RunCommand()
}

func testInsertBookList() {
	bookList := models.BookList{}
	bookList.Name = "从菜鸟到大师"
	bookList.Instruction = "职场的香饽饽大师推荐读物"
	bookList.Publish = !true
	bookList.PublishTime = time.Now()
	bookList.UserInfo = &models.User{}
	bookList.BookListType = &models.BookListType{}
	id, err := bookList.Insert()
	fmt.Println("id ", id, " ,err:", err, ", bookList ", bookList)
}

func testInsertBookInfo() {
	info := models.BookInfo{}
	info.BookType = &models.BookType{}
	info.PublishTime = time.Now()
	info.Author = "Unknown"
	info.Introduction = "null"
	info.Reward = 10
	info.SaveName = "no.pdf"
	info.Link = "http://www.github.com/xuelike/**.pdf"
	info.ContentLegal = true
	info.UserInfo = &models.User{Id: 2}
	id, err := info.Insert()
	fmt.Println("id : ", id, " , err: ", err, " , nfo ", info)
}

func testReadBookInfos()  {
	all:=make([]*models.BookInfo,0)
	orm.NewOrm().QueryTable(&models.BookInfo{}).Filter("id",2).All(&all)
	for v:=range  all{
		fmt.Printf("v = %+v\n",*all[v] )

	}
}

func testReadAllBookListTypes()  {
	all:=make([]*models.BookListType,0)
	orm.NewOrm().QueryTable(&models.BookListType{}).All(&all)
	for v:=range  all{
		fmt.Printf("BooklistType index %d , value = %+v\n",v ,*all[v] )

	}
}

// 测试多对多插入数据
func testInsertM2mBookListBookInfo()  {
	info := models.BookInfo{}
	info.UserInfo = &models.User{Id:8}
	info.BookType = &models.BookType{Id:1}
	info.PublishTime = time.Now()
	info.Link = "http://9p.cat-v.org/9p.jpg"
	info.Name = "p9"
	info.Reward = 10
	info.Introduction ="create intel os with keyboard"
	info.Author = "gitPc"
	id, err := info.Insert()
	fmt.Println(id, err )

	bookList := models.BookList{Id: 4}
	ormer := orm.NewOrm()
	// 多对多插入数据....
	ormer.QueryM2M(&bookList,"BookInfo").Add(info)

}