package controllers

import (
	"strconv"
	"gateway/models"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

//后台首页
func (this *MainController) Index() {
	this.display()
}

//登录
func (this *MainController) Login() {
	email := this.Input().Get("email")
	password := this.Input().Get("password")

	if this.IsPost(){
		if this.LoginMethod(email,password) {
			this.redirect("/index")
		}
		this.setFlash("error","账号不存在")
	}
	this.Data["email"] = email
	this.Data["password"] = password
	this.TplName = "login.html"
}

//登录方法
func (this *MainController) LoginMethod(user ...string) bool {
	email := user[0]
	password := this.crypt(user[1])
	users := new(models.User)
	u,err := users.GetUserByEmail(email,password)
	if err != nil {
		return false
	}
	this.SetSession("userId",strconv.Itoa(u.Id))
	return true
}
