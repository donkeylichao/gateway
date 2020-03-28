package routers

import (
	"gateway/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	// 跨域设置
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

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
