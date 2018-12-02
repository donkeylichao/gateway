package controllers

import (
	"github.com/astaxie/beego"
	"gateway/models"
	"strings"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	//"fmt"
)

type BaseController struct {
	beego.Controller
	user *models.User
	controllerName string
	actionName string
	userId int
	userName string
	pageSize int
}

// Prepare runs after Init before request function execution.
func (this *BaseController) Prepare()  {
	this.pageSize = 10
	controllerName,actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0:len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.auth()
	//fmt.Printf("%s",this.controllerName)
	//fmt.Printf("%s",this.actionName)

	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
}

/**
验证用户登陆
 */
func (this *BaseController) auth()  {
	userId := this.GetSession("userId")
	if userId != nil {
		users := new(models.User)
		Id,_:=strconv.Atoi(userId.(string))
		user,err := users.FindById(Id)
		if err == nil {
			this.userId = user.Id
			this.userName = user.Name
			this.user = user
		}
	}

	if this.userId == 0 && this.notLogin() {
		this.redirect(beego.URLFor("MainController.Login"))
	}
}

/**
判断是否需要跳转登录页面
 */
func (this *BaseController) notLogin() bool {
	if this.controllerName != "main" {//不是登陆页面
		return true
	}
	if this.controllerName == "main" && this.actionName != "logout" && this.actionName != "login" {
		return true
	}
	return false
}

//渲染模版
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}
	//panic(tplname)
	//fmt.Printf("s",tplname)
	this.Layout = "layout/layout.html"
	this.TplName = tplname
	this.setSessToFlash()
}


//重定向方法
func (this *BaseController) redirect(url string)  {
	this.Redirect(url,302)
	this.StopRun()
}

/**
判断是否是post请求
 */
func (this *BaseController) IsPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//md5加密
func (this *BaseController) crypt(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

//添加提示信息
func (this *BaseController) setFlash(str ...string) {
	flash := beego.NewFlash()
	if str[0] == "error" {
		flash.Error(str[1])
	}
	if str[0] == "warning" {
		flash.Warning(str[1])
	}
	if str[0] == "notice" {
		flash.Notice(str[1])
	}
	if str[0] == "success" {
		flash.Success(str[1])
	}
	flash.Store(&this.Controller)
}

/**
session 信息同步flash
 */
func (this *BaseController) setSessToFlash() {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	if success := sess.Get("success"); success != nil {
		this.setFlash("success",success.(string))
		sess.Delete("success")
	}
	if warning := sess.Get("warning");warning != nil {
		this.setFlash("warning",warning.(string))
		sess.Delete("warning")
	}
	if notice := sess.Get("notice");notice != nil {
		this.setFlash("notice",notice.(string))
		sess.Delete("notice")
	}
	if err := sess.Get("error");err != nil {
		this.setFlash("error",err.(string))
		sess.Delete("error")
	}
}

//获取客户端IP地址
func (this *BaseController) getRemoteIp() string {
	req := this.Ctx.Request
	addr := req.RemoteAddr // "IP:port" "192.168.1.150:8889"
	beego.BeeLogger.Debug("addr:%s", addr)
	return addr[1: len(addr)-4]
}

//session 存储临时信息
func (this *BaseController) SetSession(key, value interface{}) {
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	sess.Set(key.(string), value.(string))
	//fmt.Printf("%s", sess.Get(key))
}
