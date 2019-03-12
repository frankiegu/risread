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
			beego.NSRouter("/cstorage", &controllers.UserController{}, "post:CloudStorage"),
			// 查看所有上传的书籍
			beego.NSRouter("/viewbooks", &controllers.UserController{}, "get:PullCloudStorage"),
		),

		beego.NSNamespace("/books",
			// 获取最近的书单
			beego.NSRouter("/bookListAdv", &controllers.BookOps{}, "get:FetchBookListAdv"),
			// 创建新的书单
			beego.NSRouter("/establishBookList", &controllers.BookOps{}, "post:CreateBookList"),
			// 添加书籍到书单
			beego.NSRouter("/add2BookList", &controllers.BookOps{}, "post:AddBook2BookList"),
			// 创建书单的评论
			beego.NSRouter("/submitBookListComment", &controllers.BookOps{}, "post:CommitBookListComment"),
			// 获取书单的评论
			beego.NSRouter("/bookListComments", &controllers.BookOps{}, "get:FetchBookListComments"),

			// 获取书单的书籍消息 (获取的是用户公开的书籍消息)
			beego.NSRouter("/bookListDetail", &controllers.BookOps{},"get:FetchBookDetail"),
			// 获取用户自己的书单
			beego.NSRouter("/ownBookLists",&controllers.BookOps{},"get:FetchOwnBookList"),
			// 获取用户的书单
			beego.NSRouter("/userBookList", &controllers.BookOps{}, "get:FetchUserBookList"),
			beego.NSRouter("/bookList", &controllers.BookOps{}, "get:FetchBooksByBookListId"),
			// 获取某本书籍的评论
			beego.NSRouter("/bookComment", &controllers.BookOps{}, "get:FetchBookInfoComments"),
			// 创建书籍的评论
			beego.NSRouter("/commitBookComment", &controllers.BookOps{}, "post:CreateBookInfoComment"),
			// 根据书单获取书籍的列表
			beego.NSRouter("/booksBk", &controllers.BookOps{}, "get:FetchBooksByBookListId"),
			// 推荐书籍
			beego.NSRouter("/recommend", &controllers.BookOps{}, "get:FetchRecommendBooks"),
			// 上传图片资源文件
			beego.NSRouter("/dpImg", &controllers.BookOps{}, "post:DPImg"),

			// 上传书籍资料
			beego.NSRouter("/commitBookInfo",&controllers.BookOps{},"post:CommitBookInfo"),

			// 获取书单类型的消息
			beego.NSRouter("/bookListTypes",&controllers.BookOps{},"get:BookListTypes"),

			// 用户自己发布的书籍总览
			beego.NSRouter("/ownbookinfos",&controllers.BookOps{}, "get:OwnBookInfos"),
		),

	)
	beego.AddNamespace(ns)

	beego.InsertFilter("/static/upload/*", beego.BeforeStatic, func(context *context.Context) {
		log.Printf("before router filter ")
		auth := context.Input.Header(controllers.JwtAuth)

		log.Println("auth:", auth)
		// todo
	})

}
