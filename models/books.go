package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//ID, 用户的ID, 书名, 上传日期, 大小, 保存的连接名(时间戳+用户id格式化+.pdf)
type UploadBook struct {
	Id         int64     `json:"id" orm:"pk,column(id)"`
	UserInfo   *User     `orm:"rel(fk)" json:"-"`
	BookName   string    `json:"book_name"`
	UploadTime time.Time `json:"upload_time"`
	Size       int64     `json:"size"`
	SaveName   string
}

func (this *UploadBook) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

// 根据用户读取用户上传的书籍
func ReadBooksWithUser(user User) (uploadBook []*UploadBook, i int64, err error) {
	i, err = orm.NewOrm().QueryTable(UploadBook{}).Filter("user_info_id", user.Id).All(&uploadBook)
	return uploadBook, i, err
}

// 书单的类型
type BookListType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 书单
type BookList struct {
	Id           int64         `orm:"pk"`
	UserInfo     *User         `orm:"rel(fk)" json:"-"`
	BookListType *BookListType `orm:"rel(fk)"`
	Name         string
	Instruction  string // 书单的主要介绍
	PublishTime  time.Time
	Publish      bool
	BookInfo     []*BookInfo `orm:"rel(m2m);"`
}

func (this *BookList) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

// 根据用户的id 获取用户的书单
func FetchBookListByUid(uid int64) (bookLists []*BookList, err error) {
	_, err = orm.NewOrm().QueryTable(&BookList{}).Filter("user_info_id").All(&bookLists)
	return
}

// 书单中的书名
type BookListBook struct {
	Id               int64     `orm:"fk"`
	BookList         *BookList `orm:"rel(fk)"`
	Name             string
	BookListFavorite []*BookListFavorite
}

// 书单的评论消息
type BookListComment struct {
	Id          int64
	BookList    *BookList `orm:"rel(fk)"`
	UserInfo    *User     `orm:"rel(fk)"`
	Content     string
	PublishTime time.Time
	Mark        int // 评论的分数
	ScanTimes   int // 浏览次数
}

func (this *BookListComment) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

// 收藏的书单
type BookListFavorite struct {
	Id           int64
	BookList     *BookList `orm:"rel(fk)"`
	UserInfo     *User     `orm:"rel(fk)"`
	FavoriteTime time.Time
}

func (this *BookListFavorite) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

// 读书心得
type BookAttainment struct {
	Id          int64
	UserInfo    *User     `orm:"rel(fk)"`
	BookInfo    *BookInfo `orm:"rel(fk)"`
	PublishTime time.Time
	Content     string
	ScanTimes   int
}

func (this *BookAttainment) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

type BookType struct {
	Id   int64
	Name string
}

func (this *BookType) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

func (this *BookType) Read(cols ... string) (err error) {
	return orm.NewOrm().Read(this, cols ...)
}

// 书籍的基本信息
type BookInfo struct {
	Id              int64
	UserInfo        *User  `orm:"rel(fk)" json:"-"` // 上传用户
	Link            string `json:"link"`            // 书籍的连接
	Name            string `json:"name"`
	Copyright       string `json:"copyright"`
	Cover           string `json:"cover"` // 封面
	ContentLegal    bool   `json:"content_legal"`
	PublishTime     time.Time              // 发布时间
	AuditTime       time.Time              // 审核时间
	Reward          int    `json:"reward"` // 阅读完成奖励
	Author          string `json:"author"`
	DownloadTimes   int                                         // 下载次数
	Introduction    string             `json:"introduction"`    // 书籍的介绍
	BookType        *BookType          `orm:"rel(fk)" json:"-"` // 书籍的类型
	SaveName        string             `json:"save_name"`
	BookInfoComment []*BookInfoComment `orm:"reverse(many)"` // 反向多个关系
	BookLists       []*BookList        `orm:"reverse(many)"`
}

func (this *BookInfo) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

// 书籍的评论
type BookInfoComment struct {
	Id          int64     `json:"id"`
	UserInfo    *User     `orm:"rel(fk); reverse(one);" json:"user_info" `
	BookInfo    *BookInfo `orm:"rel(fk);reverse(one)" json:"-"`
	Content     string    `json:"content"`
	PublishTime time.Time `json:"publish_time"`
	ScanTimes   int       `json:"scan_times"`
}

func (this *BookInfoComment) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(this)
}

func (this *BookList) Read(cols ...string) error {
	return orm.NewOrm().Read(this, cols...)
}
