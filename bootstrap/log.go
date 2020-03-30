package bootstrap

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	//日志初始化存储位置
	logDirver := beego.AppConfig.String("log::driver")
	switch logDirver {
	case "file":
		logFile := beego.AppConfig.String("log::file")
		beego.SetLogger(logDirver, `{"filename":"`+logFile+`"}`)
	case "es":
		dsn := beego.AppConfig.String("log::dsn")
		level := beego.AppConfig.String("log::level")
		if err := beego.SetLogger(logs.AdapterEs, `{"dsn":"`+dsn+`","level":`+level+`}`); err != nil {
			logs.Error(err)
		}
		logs.Async(1e3)
	default:

	}
}
