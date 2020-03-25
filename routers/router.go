package routers

import (
	"gateway/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/index", &controllers.MainController{}, "*:Index")
	//beego.Router("/login", &controllers.MainController{}, "*:Login")
	//beego.AutoRouter(&controllers.UserController{})
	//beego.AutoRouter(&controllers.UrlController{})
	//beego.AutoRouter(&controllers.ApiController{})

	ns := beego.NewNamespace("/admin",

		beego.NSRouter("/", &controllers.MainController{}, "*:Index"),
		beego.NSRouter("/logout", &controllers.MainController{}, "*:Logout"),
		beego.NSRouter("/login", &controllers.MainController{}, "*:Login"),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/url",
			beego.NSInclude(&controllers.UrlController{}),
		),
		beego.NSNamespace("/api",
			beego.NSInclude(&controllers.ApiController{}),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/*", &controllers.EntranceController{},"*:Entrance")
}
