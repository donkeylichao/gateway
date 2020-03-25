package controllers

import (
	"github.com/astaxie/beego"
	"gateway/service/route"
	"gateway/service/http"
)

type Message struct {
	Name string
	Age int
	Phone string
}

type EntranceController struct {
	beego.Controller
}

/**
 网关入口
 */
func (c *EntranceController) Entrance() {

	c.Data["json"] = "test"
	c.ServeJSON()
	return
	requestParams := c.parseRequestParameters()

	response := make(map[string]interface{})
	response["status"] = 1
	response["success"] = "success"

	if requestParams["method"] == "OPTIONS" {
		c.Data["json"] = response
	} else {

		matchRoute := route.GetGatewayService(requestParams)
		if matchRoute == nil {
			matchRoute = "请求路由不存在"
			response["success"] = "fail"
			response["status"] = "0"
		} else {
			response,_,_ := http.Request(matchRoute.(string),requestParams)
			matchRoute = response
		}
		response["data"] = matchRoute
		c.Data["json"] = response
	}
	c.ServeJSON()
}

/**
 解析请求参数
 */
func (c *EntranceController) parseRequestParameters() map[string]interface{} {

	/* 获取请求参数 */
	method := c.Ctx.Request.Method
	header := c.Ctx.Request.Header
	body := c.Ctx.Request.Body
	form := c.Ctx.Request.Form
	multipart := c.Ctx.Request.MultipartForm
	path := c.Ctx.Request.RequestURI
	requestBody := c.Ctx.Input.RequestBody

	requestParams := make(map[string]interface{})

	requestParams["method"] = method
	requestParams["header"] = header
	requestParams["body"] = body
	requestParams["form"] = form
	requestParams["multipart"] = multipart
	requestParams["path"] = path
	requestParams["requestBody"] = requestBody

	return requestParams
}