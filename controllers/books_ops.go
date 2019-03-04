package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
	"risread/models"
	"time"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"strings"
)

type BookOps struct {
	beego.Controller
}

// 书单的评论
func (this *BookOps) CommitBookListComment() {
	dp := struct {
		UserInfoId int64  `json:"user_info_id"`
		BookListId int64  `json:"book_list_id"`
		Content    string `json:"content"`
	}{}

	var err error
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &dp)
	if err != nil {
		fmt.Println("parse CommitBookListComment err:", err)
		return
	}
	bookListComment := models.BookListComment{
		UserInfo:    &models.User{Id: dp.UserInfoId},
		BookList:    &models.BookList{Id: dp.BookListId},
		Content:     dp.Content,
		PublishTime: time.Now(),
	}

	id, err := bookListComment.Insert()
	if err != nil {
		fmt.Println("insert book list comment err:", err)
		return
	}
	fmt.Println("id :", id)

}

// 提交书籍消息
func (this *BookOps) CommitBookInfo() {
	dp := models.BookInfo{}
	authorization := this.Ctx.Input.Header("authorization")
	user, ok := gJwt[authorization]
	if !ok {
		fmt.Println("user not load ")
		this.Abort("407")
		return
	}
	dp.Reward ,_= this.GetInt("reward",-1)
	dp.Link = this.GetString("link","")
	dp.Author = this.GetString("author","")
	dp.Introduction = this.GetString("instruction","")
	dp.Name = this.GetString("name","")
	dp.Copyright = this.GetString("copyright","")
	dp.PublishTime = time.Now()
	dp.BookType =&models.BookType{Id:2}

	//err := json.Unmarshal(this.Ctx.Input.RequestBody, &dp)
	//	//
	//	//if err !=nil {
	//	//	fmt.Println("fetch data err ",err )
	//	//	this.Abort("403")
	//	//	return
	//	//}
	dp.UserInfo = &user
	//if dp.Name == "" || dp.Link == "" || dp.Reward <=0 ||dp.Reward>10 {
	//	fmt.Println("error ")
	//	this.Abort("403")
	//	return
	//}

	fmt.Printf("book info %+v \n",dp )
	// 保存封面
	cover, hCover, err := this.GetFile("cover")
	if err != nil {
		fmt.Println("get file err ", err)
		this.Abort("401")
		return
	}
	defer cover.Close()

	hash := md5.New()
	hash.Write([]byte(time.Now().String()))
	bytes := hash.Sum(nil)
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "%x", bytes)
	fmt.Println("s ", builder.String())
	splits := strings.Split(hCover.Filename, ".")
	//
	if len(splits) != 2 {
		fmt.Println("file error : ", "len is not 2 ",hCover.Filename)
		this.Abort("403")
		return
	}
	CoverUrl :="static/img/" + builder.String() + "." + splits[1]
	// dp 保存数据信息

	// 保存图片文件..
	err = this.SaveToFile("cover", CoverUrl) // 保存位置在 static/upload, 没有文件夹要先创建
	if err !=nil {
		fmt.Println("save file error ",err)
		this.Abort("403")
		return
	}

	pdfFile, pdfHeader, err := this.GetFile("pdf")
	if err !=nil {
		fmt.Println("err: ",err)
		this.Abort("403")
		return
	}
	defer pdfFile.Close()

	hash.Reset()
	hash.Write([]byte(time.Now().String()))
	bytes = hash.Sum(nil )
	builder.Reset()
	fmt.Fprintf(&builder,"%x",bytes)

	splits = strings.Split(pdfHeader.Filename, ".")
	//
	if len(splits) != 2 {
		fmt.Println("file error : ", "len is not 2 ")
		this.Abort("403")
		return
	}

	pdfUrl :="static/upload/" + builder.String() + "." + splits[1]
	// 保存 上传的文件资源
	err=this.SaveToFile("pdf",pdfUrl)
	if err !=nil {
		fmt.Println("err save pdf err ")
		this.Abort("403")
		return
	}

	dp.SaveName = pdfUrl
	dp.Cover = CoverUrl
	id, err := dp.Insert()
	if err !=nil {
		fmt.Println("err:",err)
		this.Abort("403")
		return
	}

	fmt.Println("ok ",dp )
	dp.Id = id
	this.Data["json"] = dp
	this.ServeJSON(true )

}

// 获取书单的评论
func (this *BookOps) FetchBookListComments() {

}

// 获取用户的书单
func (this *BookOps) FetchUserBookList() {
	uid, err := this.GetInt64("uid", -1)
	if err != nil {
		fmt.Println("getUid err:", err)
		return
	}
	//bookLists :=make([]*models.BookList,0)
	bookLists, err := models.FetchBookListByUid(uid)
	if err != nil {
		fmt.Println("fetchBookError ", err)
		return
	}

	for _, v := range bookLists {
		fmt.Printf("value =%+v\n", v)
	}
	this.Data["json"] = bookLists
	this.ServeJSON(true)
	return

}

// 获取书单的详细消息
func (this *BookOps) FetchBookDetail() {
	// 获取书单的id
	listId, err := this.GetInt64("listId", -1)
	if err != nil {
		fmt.Println("getBookListId err ", err)
		return
	}

	fmt.Println("listId ", listId)
}

// 添加书籍到书单
func (this *BookOps) AddBook2BookList() {

	dp := struct {
		BookListId int64 `json:"book_list_id"`
		BookInfoId int64 `json:"book_info_id"`
	}{}
	fmt.Println(dp)
}

// 创建书单
func (this *BookOps) CreateBookList() {
	dp := struct {
		UserInfoId  int64  `json:"user_info_id"`
		Instruction string `json:"instruction"`
		Name        string `json:"name"`
		Publish     bool   `json:"publish"`
		TypeId      int    `json:"type_id"`
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dp)
	if err != nil {
		fmt.Println("parse CreateBookList err:", err)
		return
	}
	authorization := this.Ctx.Input.Header("authorization")
	user, ok := gJwt[authorization]
	if !ok {
		this.Abort("404")
		return
	}
	bookList := &models.BookList{}
	bookList.PublishTime = time.Now()
	bookList.UserInfo = &models.User{Id: user.Id}
	bookList.Instruction = dp.Instruction
	bookList.Name = dp.Name
	bookList.Publish = dp.Publish
	bookList.BookListType = &models.BookListType{Id: dp.TypeId}
	id, err := bookList.Insert()
	if err != nil {
		fmt.Println("insert bookListError err:", err)
		return
	}
	fmt.Println("id ", id, " bookList ", bookList)

	this.Data["json"] = struct {
		Code       int
		BookListId int64
	}{
		Code:       200,
		BookListId: id,
	}
	this.ServeJSON(true)
	return
}

//  获取书单
func (this *BookOps) FetchBookListAdv() {

}

// 获取自己的书单消息
func (this *BookOps) FetchOwnBookList() {
	authorization := this.Ctx.Input.Header("authorization")

	user, ok := gJwt[authorization]
	if !ok {
		this.Abort("404")
		return
	}
	dp := []struct {
		Name        string    `json:"name"`
		Id          int64     `json:"id"`
		Instruction string    `json:"instruction"`
		PublishTime time.Time `json:"publish_time"`
		Publish     bool      `json:"publish"`
	}{}

	// 查找用户自己的书单
	sql := "SELECT T0.id , T0.name , T0.instruction , T0.publish_time , T0.publish " +
		"FROM book_list AS T0 WHERE" +
		" T0.user_info_id = ?";
	n, err := orm.NewOrm().Raw(sql, user.Id).QueryRows(&dp)

	if err != nil {
		fmt.Println("err ", err)
		this.Abort("404")
		return
	}
	fmt.Println("read n ", n)
	this.Data["json"] = dp
	this.ServeJSON(true)
	return
}

// 添加书籍的评论
func (this *BookOps) CreateBookInfoComment() {
	dp := struct {
		BookInfoId int64  `json:"book_info_id"`
		Content    string `json:"content"`
		UserInfoId int64  `json:"user_info_id"`
	}{}
	fmt.Println(dp)
	var err error
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &dp)
	if err != nil {
		fmt.Println("parse CreateBookInfoComment err ", err)
		return
	}

	bookInfoComment := models.BookInfoComment{
		UserInfo:    &models.User{Id: dp.UserInfoId},
		BookInfo:    &models.BookInfo{Id: dp.BookInfoId},
		Content:     dp.Content,
		PublishTime: time.Now(),
	}

	id, err := bookInfoComment.Insert()
	if err != nil {
		fmt.Println("bookInfoComment insert err:", err)
		return
	}

	fmt.Println("id ", id, "bookInfoComment ", bookInfoComment)

}

// 获取某本书籍的评论
func (this *BookOps) FetchBookInfoComments() {
	bookCID, err := this.GetInt64("bookInfoComId", -1)
	if err != nil {
		fmt.Println("getBookCID err: ", err)
		return
	}

	fmt.Println("cid ", bookCID)
	offset, err := this.GetInt64("offset", -1)

	if offset == -1 {

	}

	//var bookInfoComments []*models.BookInfoComment

	var bookInfoComments []*struct {
		Id           int64 `json:"id"`
		Content      string
		ScanTimes    int `json:"scan_times"`
		PublishTime  time.Time
		Author       string
		Username     string
		Link         string
		Introduction string
		UserInfoId   int64 `json:"user_info_id"`
		BookInfoId   int64 `json:"book_info_id"`
	}

	n, err := orm.NewOrm().Raw(
		"SELECT T0.id ,"+
			"T0.content,"+
			"T0.scan_times ,"+
			"T0.publish_time,"+
			"T1.id  as user_info_id ,T1.username,"+
			"T2.id as book_info_id,T2.link,"+
			"T2.content_legal,"+
			"T2.publish_time,"+
			"T2.author,"+
			"T2.download_times,"+
			"T2.introduction "+
			"From "+
			"book_info_comment T0 ,"+
			" user_info T1 , "+
			"book_info T2 WHERE"+
			" T0.book_info_id = T2.id AND T0.user_info_id = T1.id  AND  T2.id = ?"+
			" ORDER BY T0.publish_time DESC", bookCID).
		QueryRows(&bookInfoComments, )

	if err != nil {
		fmt.Println("error ", err, n)
		return
	}
	fmt.Println(n)
	this.Data["json"] = bookInfoComments
	this.ServeJSON(true)
}

// 提交书籍的评论消息
func (this *BookOps) CommitBookInfoComment() {
	dp := struct {
		UserInfoId int    `json:"user_info_id"`
		BookInfoId int    `json:"book_info_id"`
		Content    string `json:"content"`
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dp)

	if err != nil {
		fmt.Println("parse CommitBookInfoComment err: ", err)
		return
	}

	//	 持久化
	bookInfoComment := models.BookInfoComment{}
	bookInfoComment.UserInfo = &models.User{Id: int64(dp.UserInfoId)}
	bookInfoComment.BookInfo = &models.BookInfo{Id: int64(dp.BookInfoId)}
	bookInfoComment.Content = dp.Content
	bookInfoComment.PublishTime = time.Now()

	id, err := bookInfoComment.Insert()
	if err != nil {
		fmt.Println("insertBookInfoComment err: ", err)
		return
	}
	fmt.Println(bookInfoComment.Id, id)
	return
}

// 根据书单获取书籍list
func (this *BookOps) FetchBooksByBookListId() {
	bkId, err := this.GetInt64("bkId", -1)
	if err != nil {
		fmt.Println("get bkId err: ", err)
		return
	}
	var bk [] struct {
		Name      string `json:"name"`
		Link      string `json:"link"`
		Author    string `json:"author"`
		Id        int64  `json:"id"`
		Copyright string `json:"copyright"`
		Cover     string `json:"cover"`
		SaveName  string `json:"save_name"`
	}
	sql := "SELECT T0.link ,T0.author ,T0.name , T0.copyright ,T0.cover," +
		"T0.id AS id " +
		"FROM book_info T0 , book_list T1 , book_list_book_infos T2  " +
		"WHERE T0.id = T2.book_info_id AND T1.id = t2.book_list_id AND T1.id = ?"
	i, e := orm.NewOrm().Raw(sql, bkId).QueryRows(&bk)
	fmt.Println(i, e)
	this.Data["json"] = bk
	this.ServeJSON(true)

}

// 获取推荐的书籍
func (this *BookOps) FetchRecommendBooks() {
	typeId, err := this.GetInt64("typeId", -1)
	if err != nil {
		this.Abort("403")
	}
	fmt.Println("typeName ", typeId)

	var recommends []*models.BookInfo

	i, err := orm.NewOrm().Raw(
		`SELECT T0.id ,
T0.link ,
T0.author ,
T0.audit_time, 
T0.download_times ,
T0.publish_time,
T0.save_name,
T0.reward 
FROM book_info T0 where T0.book_type_id = ? order by publish_time DESC LIMIT 10`, typeId).QueryRows(&recommends)
	fmt.Println(i, err)

	this.Data["json"] = recommends
	this.ServeJSON(true)
}

//用户上传图片
func (this *BookOps) DPImg() {
	f, h, err := this.GetFile("img")
	if err != nil {
		fmt.Println("getfile err ", err)
		this.Abort("401")
		return
	}
	defer f.Close()

	hash := md5.New()
	hash.Write([]byte(time.Now().String()))
	bytes := hash.Sum(nil)
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "%x", bytes)
	fmt.Println("s ", builder.String())
	splits := strings.Split(h.Filename, ".")
	//
	if len(splits) != 2 {
		fmt.Println("file error : ", "len is not 2 ")
		this.Abort("403")
		return
	}

	// dp 保存数据信息
	dp := struct {
		Cover string `json:"cover"`
	}{
		Cover: "static/img/" + builder.String() + "." + splits[1],
	}
	// 保存文件..
	this.SaveToFile("img", dp.Cover) // 保存位置在 static/upload, 没有文件夹要先创建
	this.Data ["json"] = dp
	this.ServeJSON(true)
}
