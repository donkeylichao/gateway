package help

import (
	"github.com/astaxie/beego"
	"fmt"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
)

var Redis cache.Cache

func init() {
	collectionName := beego.AppConfig.String("redis::name")
	host := beego.AppConfig.String("redis::host")
	port := beego.AppConfig.String("redis::port")
	db := beego.AppConfig.String("redis::db")
	password := beego.AppConfig.String("redis::password")

	redisCon := `{"key":"` + collectionName + `","conn":"` + host + `:` + port + `","dbNum":"` + db + `","password":"` + password + `"}`

	redis, err := cache.NewCache("redis", redisCon)

	if err == nil {
		fmt.Println("链接redis成功")
		Redis = redis
		return
	}
	fmt.Println(err.Error())
}
