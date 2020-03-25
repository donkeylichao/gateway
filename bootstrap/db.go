package bootstrap

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("mysql::host")
	dbport := beego.AppConfig.String("mysql::port")
	dbuser := beego.AppConfig.String("mysql::user")
	dbpassword := beego.AppConfig.String("mysql::password")
	dbname := beego.AppConfig.String("mysql::name")
	dbtimezone := beego.AppConfig.String("mysql::timezone")

	if dbport == "" {
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	if dbtimezone != "" {
		dsn = dsn + "&loc" + url.QueryEscape(dbtimezone)
	}

	orm.RegisterDataBase("default", "mysql", dsn)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
