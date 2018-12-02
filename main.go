package main

import (
	_ "gateway/routers"
	"github.com/astaxie/beego"
	_ "gateway/models"
	_ "gateway/help"
)

func main() {
	beego.Run()
}
