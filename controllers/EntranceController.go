package controllers

import (
	"github.com/astaxie/beego"
	"gateway/service/route"
	"strings"
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

	requestParams := c.parseRequestParameters()

	response := make(map[string]interface{})

	if requestParams["method"] == "OPTIONS" {
		response["status"] = 1
		response["msg"] = "success"
	} else {

		matchRoute := route.GetGatewayService(requestParams)
		apiReturn := map[string]interface{}{}

		if matchRoute == "" {
			response["status"] = 0
			response["msg"] = "请求路由不存在"
		} else {
			responseData, err := http.Request(requestParams,matchRoute)
			if err != nil {
				response["status"] = 0
				response["msg"] = err.Error()
			} else {
				response["status"] = 1
				response["msg"] = "success"
			}
			apiReturn = responseData
		}
		response["data"] = apiReturn
	}

	c.Data["json"] = response
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
	allpath := c.Ctx.Request.RequestURI
	requestBody := c.Ctx.Input.RequestBody
	index := strings.Index(allpath,"?")
	path := ""
	query := ""
	if index < 0 {
		path = allpath
	} else {
		path = allpath[0:index]
		query = allpath[index:]
	}

	requestParams := make(map[string]interface{})

	requestParams["method"] = method
	requestParams["header"] = header
	requestParams["body"] = body
	requestParams["form"] = form
	requestParams["multipart"] = multipart
	requestParams["path"] = path
	requestParams["query"] = query
	requestParams["requestBody"] = requestBody

	return requestParams
}