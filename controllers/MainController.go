package controllers

import (
	"strconv"
	"gateway/models"
	"github.com/astaxie/beego"
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
func (c *MainController) Index() {
	c.display()
}

//登录
func (c *MainController) Login() {
	email := c.Input().Get("email")
	password := c.Input().Get("password")

	if c.IsPost(){
		if c.LoginMethod(email,password) {
			c.redirect(beego.URLFor("MainController.Index"))
		}
		c.setFlash("error","账号不存在")
	}
	c.Data["email"] = email
	c.Data["password"] = password
	c.TplName = "login.html"
}

//登录方法
func (c *MainController) LoginMethod(user ...string) bool {
	email := user[0]
	password := c.crypt(user[1])
	users := new(models.User)
	u,err := users.GetUserByEmail(email,password)
	if err != nil {
		return false
	}
	c.SetSession("userId",strconv.Itoa(u.Id))
	return true
}

//退出
func (c *MainController) Logout() {
	c.DelSession("userId")
	c.redirect(beego.URLFor("MainController.Login"))
}
