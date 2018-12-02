package routers

import (
	"gateway/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
    beego.AutoRouter(&controllers.UserController{})
    beego.AutoRouter(&controllers.UrlController{})
    beego.AutoRouter(&controllers.ApiController{})
}
