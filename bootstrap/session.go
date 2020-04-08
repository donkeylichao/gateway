package bootstrap

import "github.com/astaxie/beego"

func init() {
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
}
