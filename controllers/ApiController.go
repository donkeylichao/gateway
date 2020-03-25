package controllers

import (
	"gateway/models"
	"gateway/help"
	"github.com/astaxie/beego"
	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type ApiController struct {
	BaseController
}

/**
条件索搜列表
 */
func (c *ApiController) List() {
	var api models.ServiceApi
	page, _ := c.GetInt("page")
	serviceName := c.GetString("name")
	method := c.GetString("method")

	condition := make([]interface{}, 0)
	if serviceName != "" {
		condition = append(condition, "service_name")
		condition = append(condition, serviceName)
	}
	if method != "" {
		condition = append(condition, "method")
		condition = append(condition, method)
	}
	condition = append(condition, []interface{}{"is_delete", models.IS_DELETE_NO}...)
	list, count := api.ConditionList(page, condition...)

	urlModel := []models.ServiceUrl{}
	_, err := orm.NewOrm().QueryTable(new(models.ServiceUrl)).Filter("is_delete", models.IS_DELETE_NO).All(&urlModel)
	if err != nil {
		logs.Error(err)
	} else {
		api.AddServiceName(list, urlModel)
	}

	c.Data["list"] = list
	c.Data["pageTitle"] = "api列表"
	c.Data["pageBar"] = help.NewPager(page, int(count), c.pageSize, beego.URLFor("ApiController.List"), true).ToString()
	c.display()

	help.Redis.Delete(beego.AppConfig.String("route.cache"))
}

/**
添加api
 */
func (c *ApiController) Create() {
	var api models.ServiceApi
	if c.IsPost() {
		//api.ServiceName = strings.TrimSpace(c.GetString("service_name"))
		//api.Method = strings.TrimSpace(c.GetString("method"))
		//api.ApiAlias = strings.TrimSpace(c.GetString("api_alias"))
		//api.ApiPath = strings.TrimSpace(c.GetString("api_path"))
		//api.InnerPath = strings.TrimSpace(c.GetString("inner_path"))
		if err := c.ParseForm(&api); err == nil {
			if _, err := govalidator.ValidateStruct(api); err == nil {
				_, err := api.Create()
				if err == nil {
					c.SetSession("success", "添加成功")
					c.redirect(beego.URLFor("ApiController.List"))
				}
				c.setFlash("error", err.Error())
			} else {
				c.setFlash("notice", err.Error())
			}
		} else {
			c.setFlash("notice", err.Error())
		}
	}

	var url []models.ServiceUrl
	o := orm.NewOrm()
	_,err := o.QueryTable(new(models.ServiceUrl)).Filter("is_delete",models.IS_DELETE_NO).All(&url)
	if err != nil{
		c.setFlash("error",err.Error())
	} else {
		c.Data["serviceUrl"] = url
	}

	c.Data["method"] = api.GetMethodAll()
	c.Data["api"] = api
	c.Data["pageTitle"] = "添加API"
	c.display()

	help.Redis.Delete(beego.AppConfig.String("route.cache"))
}

/**
删除api
 */
func (c *ApiController) Delete() {
	var api models.ServiceApi
	id, _ := c.GetInt("id");
	if _, err := api.FindById(id); err != nil {
		c.SetSession("notice", "删除数据不存在")
		c.redirect(beego.URLFor("ApiController.List"))
	}

	if _, err := api.Delete(); err == nil {
		c.SetSession("success", "删除成功")
		c.redirect(beego.URLFor("ApiController.List"))
	}
	help.Redis.Delete(beego.AppConfig.String("route.cache"))
	c.SetSession("error", "删除失败")
	c.redirect(beego.URLFor("ApiController.List"))
}

/**
修改api
 */
func (c *ApiController) Update() {
	var api models.ServiceApi
	id,_ := c.GetInt("id")
	if _,err := api.FindById(id);err != nil {
		c.SetSession("error","修改数据不存在")
		c.redirect(beego.URLFor("ApiController.List"))
	}
	if c.IsPost() {
		if err := c.ParseForm(&api);err == nil {
			if _,err := govalidator.ValidateStruct(api);err==nil {
				_,err := api.Update()
				if err ==nil {
					c.SetSession("success","修改成功")
					c.redirect(beego.URLFor("ApiController.List"))
				}
				c.setFlash("error",err.Error())
			} else {
				c.setFlash("notice",err.Error())
			}
		} else {
			c.setFlash("notice",err.Error())
		}
	}

	var url []models.ServiceUrl
	o := orm.NewOrm()
	_,err := o.QueryTable(new(models.ServiceUrl)).Filter("is_delete",models.IS_DELETE_NO).All(&url)
	if err != nil{
		c.setFlash("error",err.Error())
	} else {
		c.Data["serviceUrl"] = url
	}

	c.Data["method"] = api.GetMethodAll()
	c.Data["api"] = api
	c.Data["pageTitle"] = "api修改"
	c.display()
	help.Redis.Delete(beego.AppConfig.String("route.cache"))
}
