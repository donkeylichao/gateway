package main

import (
	_ "gateway/routers"
	"github.com/astaxie/beego"
	_ "gateway/bootstrap"
)

func main() {
	beego.Run()
}
