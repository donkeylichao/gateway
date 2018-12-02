package controllers

import (
	"github.com/astaxie/beego"
	"gateway/models"
	"gateway/help"
	"strings"
	"time"
	"github.com/asaskevich/govalidator"
)

type UserController struct {
	BaseController
}

/**
用户列表
 */
func (c *UserController) List() {
	var User models.User
	page, _ := c.GetInt("page", 0)
	username := c.GetString("username")

	condition := make([]string, 0)
	if username != "" {
		condition = append(condition, "name")
		condition = append(condition, username)
	}
	list, count := User.ConditionList(page, condition)
	c.Data["list"] = list
	c.Data["pageTitle"] = "用户列表"
	c.Data["pageBar"] = help.NewPager(page, int(count), c.pageSize, beego.URLFor("UserController.List"), true).ToString()
	c.display()
}

/**
添加用户
 */
func (c *UserController) Create() {
	var user models.User
	if c.IsPost() {
		user.Password = c.crypt(strings.TrimSpace(c.GetString("password")))
		user.LastLoginIp = "192.149.1.22"
		user.LastLoginTime = int(time.Now().Unix())

		if err := c.ParseForm(&user); err == nil {
			if _, err := govalidator.ValidateStruct(user); err == nil {
				count, err := user.Create()
				if count != 0 {
					c.SetSession("success", "添加成功")
					c.redirect(beego.URLFor("UserController.List"))
				}
				c.setFlash("error", err.Error())
			} else {
				c.setFlash("notice", err.Error())
			}
		} else {
			c.setFlash("notice", err.Error())
		}
	}
	c.Data["user"] = user
	c.Data["pageTitle"] = "添加用户"
	c.display()
}

/**
删除用户
 */
func (c *UserController) Delete() {
	var User models.User
	id, _ := c.GetInt("id")
	if _, err := User.FindById(id); err != nil {
		c.SetSession("error", "用户不存在")
		c.redirect(beego.URLFor("UserController.List"))
	}
	if _, err := User.Delete(id); err == nil {
		c.SetSession("success", "删除成功")
		c.redirect(beego.URLFor("UserController.List"))
	}
	c.SetSession("error", "删除失败")
	c.redirect(beego.URLFor("UserController.List"))
}

/**
更新用户
 */
func (c *UserController) Update() {
	var user models.User
	id, _ := c.GetInt("id", 0)
	user.FindById(id)
	//if err,_ := user.FindById(id);err == nil {
	//	c.SetSession("error","修改用户不存在")
	//	c.redirect(beego.URLFor("UserController.List"))
	//}
	//fmt.Printf("%s",user)
	if c.IsPost() {
		user.Name = c.GetString("name")
		user.Email = c.GetString("email")
		user.RealName = c.GetString("real_name")
		if c.GetString("password") != "" {
			user.Password = c.crypt(strings.TrimSpace(c.GetString("password")))
		}
		user.Phone = strings.TrimSpace(c.GetString("phone"))
		user.LastLoginTime = int(time.Now().Unix())
		user.LastLoginIp = "0.0.0.0"

		if _, err := govalidator.ValidateStruct(user); err == nil {
			count, err := user.Update()
			if count != 0 {
				c.SetSession("success", "修改成功")
				c.redirect(beego.URLFor("UserController.List"))
			}
			c.setFlash("error", err.Error())
		} else {
			c.setFlash("notice", err.Error())
		}
	}

	c.Data["user"] = user
	c.Data["pageTitle"] = "编辑用户"
	c.display()
}
