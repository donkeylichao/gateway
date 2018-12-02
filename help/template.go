package help

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/astaxie/beego/context"
)

//时间
func Date(ti int) string {
	t := time.Unix(int64(ti), 0)
	return t.Format("2006-01-02 15:04:05")
}

//序号加1模版方法
func Index(index int) int {
	return index + 1
}

//传入context 获取session
func Msg() string {
	ctx := context.NewContext()
	globalSessions := beego.GlobalSessions
	go globalSessions.GC()
	//fmt.Printf("%s", ctx.ResponseWriter)
	//return ""
	sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	//fmt.Printf("%s", sess)
	//return ""
	success := sess.Get("success")
	error := sess.Get("error")
	notice := sess.Get("notice")
	sess.Flush() //清除所有的session
	if success != "" {
		return "<div class='alert alert-success alert-dismissible' role='alert'><button type='button' class='close' data-dismiss='alert' aria-label='Close'><span>&times;</span></button><strong>" + success.(string) + "</strong></div>"
	} else if error != "" {
		return "<div class = 'alert alert-danger alert-dismissible' role = 'alert'><button type = 'button' class = 'close' data-dismiss = 'alert' aria-label = 'Close'><span>&times; </span></button><strong>" + error.(string) + "</strong></div>"
	} else if notice != "" {
		return "<div class = 'alert alert-warning alert-dismissible' role = 'alert'><button type = 'button' class = 'close' data-dismiss = 'alert' aria-label = 'Close'><span>&times; </span></button><strong>" + notice.(string) + "</strong></div>"
	}
	return ""
}

//初始化自定义的模版函数
func init() {
	beego.AddFuncMap("Date", Date)
	beego.AddFuncMap("In", Index)
	beego.AddFuncMap("Msg", Msg)
}
