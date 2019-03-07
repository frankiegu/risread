package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id       int64
	Username string
	Phone    string `json:"-"`
	Password string `json:"-"`
	Gender   string `orm:default("不公开")`
	//Address string
	Email     string `json:"-"`
	Riches    int // 用户积分
	Signature string
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func (this *User) Read(cols ... string) (err error) {
	orm := orm.NewOrm()
	return orm.Read(this, cols...)

}

func (this *User) Insert() (int64, error) {
	orm := orm.NewOrm()
	return orm.Insert(this)
}

func (this *User) TableName() string {
	return "user_info"
}
