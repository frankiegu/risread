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
func ReadBooksWithUser(user User)(uploadBook[]*UploadBook,i int64,err error )  {
	i, err = orm.NewOrm().QueryTable(UploadBook{}).Filter("user_info_id", user.Id).All(&uploadBook)
	return uploadBook,i,err
}