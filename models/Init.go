package models

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	dbtimezone := beego.AppConfig.String("db.timezone")

	if dbport == "" {
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	if dbtimezone != "" {
		dsn = dsn + "&loc" + url.QueryEscape(dbtimezone)
	}
	orm.RegisterModelWithPrefix(beego.AppConfig.String("db.prefix"),new(User),new(serviceUrl),new(serviceApi))
	orm.RegisterDataBase("default", "mysql", dsn)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
