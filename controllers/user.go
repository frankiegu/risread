package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"encoding/json"
	"regexp"
	"fmt"
	"risread/models"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/satori/go.uuid"
	"strings"
	"strconv"
)

const (
	//  手机号码格式
	PhonePattern = `^(1[3|4|5|8|6][0-9]\d{8,8})$`
	JwtAuth      = `Authorization`
)

var (
	jwtSigningKey = []byte("bla bla bla")

	errmsg = struct {
		Code int `json:"code"`
		Data interface{}
	}{
		Code: 400,
		Data: "未知错误",
	}
	// 全局jwt 列表
	gJwt = make(map[string]models.User)
)
// Operations about Users
type UserController struct {
	beego.Controller
}

// 登录
func (u *UserController) Login() {
	lp := struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{}

	// 反序列化请求数据
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &lp)
	if err != nil {
		log.Println("error ", err)
		resData(&u.Controller, nil)
		return
	}

	ok, _ := RegexpValidPattern(lp.Phone, PhonePattern)

	if !ok {

		log.Println("phone err: ")
		resData(&u.Controller, errmsg)
		return
	}

	// 校验密码是否正确
	user := models.User{
		Phone:    lp.Phone,
		Password: lp.Password,
	}

	err = user.Read("phone", "password")
	if err != nil {
		log.Println("error : ", err)
		resData(&u.Controller, errmsg)
		return
	}

	log.Printf("user :%+v\n", user)

	// jwt 返回
	// 生成uuid
	uuids, err := uuid.NewV4()
	if err != nil {
		log.Println("jenerated uuid err: ", err)
		resData(&u.Controller, errmsg)
	}

	jswt := generateJWT(uuids)

	// 设置 auth  响应头消息
	u.Ctx.ResponseWriter.Header().Set("Authorization", jswt)

	appendJwt(jswt, user)
}

func (u *UserController) Logout() {
	auth := u.Ctx.Input.Header("Authorization")
	delete(gJwt, auth)
}

// 注册
func (u *UserController) Register() {
	register := struct {
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		RePassword string `json:"re_password"`
	}{}

	err := json.Unmarshal(u.Ctx.Input.RequestBody, &register)
	if err != nil {
		log.Println(err)
		resData(&u.Controller, errmsg)
		return
	}
	// 验证手机号码是否符合格式
	ok, err := RegexpValidPattern(register.Phone, PhonePattern)
	fmt.Println(ok, err)
	if !ok {
		resData(&u.Controller, errmsg)
		return
	}
	if register.Password != register.RePassword || len(register.Password) < 8 || len(register.Password) > 16 {
		resData(&u.Controller, errmsg)
		log.Println("密码的长度不正确,请重新设置")
		return
	}
	log.Printf("register %+v\n", register)

	// 密码校验
	user := models.User{Phone: register.Phone, Password: register.Password}
	err = user.Read("phone")
	if err != orm.ErrNoRows {
		log.Println("该用户已经被注册... :", err)
		resData(&u.Controller, errmsg)
		return
	}

	user.Gender = "不公开"
	id, err := user.Insert()
	if err != nil {
		fmt.Println("持久化数据库失败", err, id)
		resData(&u.Controller, errmsg)
		return
	}

	log.Println("user: ", user)

	// jwt 返回
	// 生成uuid
	uuids, err := uuid.NewV4()
	if err != nil {
		log.Println("jenerated uuid err: ", err)
		resData(&u.Controller, errmsg)
	}

	jswt := generateJWT(uuids)

	// 设置 auth  响应头消息
	u.Ctx.ResponseWriter.Header().Set("Authorization", jswt)

	appendJwt(jswt, user)
	success := struct {
		Code    int
		Message string
	}{
		Code:    200,
		Message: "注册成功",
	}
	resData(&u.Controller, success)
}

// 上传文件到云
func (this *UserController) CloudStorage() {

	user := ValidUser(this)
	log.Println("user:", user)
	file, h, err := this.GetFile("uploadname")
	if err != nil {
		log.Println("uploadFile err: ", err)
		resData(&this.Controller, errmsg)
		return
	}

	defer file.Close()
	if !strings.HasSuffix(h.Filename, ".pdf") {
		log.Println(" upload is not  *.pdf")
		resData(&this.Controller, errmsg)
		return
	}

	book := models.UploadBook{}
	book.UserInfo = &user
	book.BookName = h.Filename
	book.Size = h.Size
	book.UploadTime = time.Now()
	book.SaveName = fmt.Sprintf("%s%d.pdf", strconv.FormatInt(time.Now().Unix(), 10), user.Id)

	err = this.SaveToFile("uploadname", "static/upload/"+book.SaveName) // 保存位置在 static/upload, 没有文件夹要先创建

	if err != nil {
		log.Println("save file err: ", err)
		resData(&this.Controller, errmsg)
		return
	}
	_, err = book.Insert()

	if err != nil {
		log.Println("cs insert upload book err:", err)
		resData(&this.Controller, errmsg)
		return
	}

	// 上传书籍成功
	resData(&this.Controller, struct {
		Code    int
		Message interface{}
	}{
		Code:    200,
		Message: "操作成功",
	})

}

// 获取上传的文件名列表
func (this *UserController) PullCloudStorage() {
	user := ValidUser(this)

	uploadBook, i, err := models.ReadBooksWithUser(user)
	log.Printf("books %+v , all +%v err %+v user %+v\n", uploadBook, i, err, user)
	resData(&this.Controller, uploadBook)

}

func ValidUser(this *UserController) (user models.User) {

	auth := this.Ctx.Input.Header(JwtAuth)
	user, ok := gJwt[auth]
	// 获取用户登录的凭证失败
	if !ok {
		log.Println("err:该用户没有登录")
		resData(&this.Controller, struct {
			Message string
			Code    int
		}{
			Message: "该用户没有登录",
			Code:    400,
		})

	}

	log.Println("user:", user)
	if !ok {
		this.StopRun()
	}
	return user

}

// 正则 检验

// 返回值
// true 表示验证通过, false 表示验证失败
// err 如果遇到错误,则返回error
func RegexpValidPattern(data string, pattern string) (ok bool, err error) {
	ok, err = regexp.MatchString(pattern, data)
	return ok, err
}

// 返回json格式的数据类型
func resData(c *beego.Controller, v interface{}) {
	c.Data["json"] = v
	c.ServeJSON(true)

}

// 生成jwt
func generateJWT(uuids uuid.UUID) (jswt string) {

	var err error
	claims := jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		// 过期时间设置为一年
		ExpiresAt: int64(time.Now().Unix() + int64(time.Hour)*24*30*12),
		Issuer:    "reading",
		Subject:   uuids.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jswt, err = token.SignedString(jwtSigningKey)
	if err != nil {
		fmt.Println("err > ", err.Error())
		return jswt
	}

	return jswt
}

// 更新jwt
func updateJwt(newjswt, oldjwt string) {
	va, ok := gJwt[oldjwt]
	if ok {
		delete(gJwt, oldjwt)

	}
	gJwt[newjswt] = va
}

// 添加新的用户
func appendJwt(jswt string, user models.User) {

	gJwt[jswt] = user
	log.Println(len(gJwt))
}
