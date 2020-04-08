package controllers

import (
	"gateway/models"
	"gateway/help"
	"github.com/astaxie/beego"
	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type UrlController struct {
	BaseController
}

/**
条件搜索列表
 */
// @router /list [get]
func (c *UrlController) List() {
	var url models.ServiceUrl
	page, _ := c.GetInt("page")
	serviceName := c.GetString("name")

	condition := make([]interface{}, 0)
	if serviceName != "" {
		condition = append(condition, "service_name")
		condition = append(condition, serviceName)
	}
	list, count := url.ConditionList(page, condition...)
	c.Data["list"] = list
	c.Data["pageTitle"] = "url列表"
	c.Data["pageBar"] = help.NewPager(page, int(count), c.pageSize, beego.URLFor("UrlController.List"), true).ToString()
	c.display()
}

/**
添加url
 */
// @router /create [get,post]
func (c *UrlController) Create() {
	var url models.ServiceUrl
	filterUrls := []string{}
	if c.IsPost() {
		urls := c.GetStrings("service_url")
		for _,v := range urls {
			if v != "" {
				filterUrls = append(filterUrls,v)
			}
		}

		if err := c.ParseForm(&url); err == nil {
			if _, err := govalidator.ValidateStruct(url); err == nil {

				if len(filterUrls) == 0 {
					c.setFlash("error", "节点不能为空")

				} else {
					serviceUrl,_ := json.Marshal(filterUrls)
					url.ServiceUrl = string(serviceUrl)
					_, err := url.Create()
					if err == nil {
						c.SetSession("success", "添加成功")
						help.Redis.Delete(beego.AppConfig.String("route.cache"))
						c.redirect(beego.URLFor("UrlController.List"))
					}
					c.setFlash("error", err.Error())
				}
			} else {
				c.setFlash("notice", err.Error())
			}
		} else {
			c.setFlash("notice", err.Error())
		}
	}

	c.Data["urls"] = filterUrls
	c.Data["urlLength"] = len(filterUrls)
	c.Data["url"] = url
	c.Data["pageTitle"] = "添加URL"
	c.display()
}

/**
删除url
 */
// @router /delete [get]
func (c *UrlController) Delete() {
	var url models.ServiceUrl
	id, _ := c.GetInt("id")
	if _, err := url.FindById(id); err != nil {
		c.SetSession("notice", "删除数据不存在")
		c.redirect(beego.URLFor("UrlController.List"))
	}
	count, _ := orm.NewOrm().QueryTable(new(models.ServiceApi)).Filter("service_url_id", id).Count()
	if count > 0 {
		c.SetSession("error", "请先删除api")
		c.redirect(beego.URLFor("UrlController.List"))
	}
	if _, err := url.Delete(); err == nil {
		c.SetSession("success", "删除成功")
		help.Redis.Delete(beego.AppConfig.String("route.cache"))
		c.redirect(beego.URLFor("UrlController.List"))
	}
	c.SetSession("error", "删除失败")
	c.redirect(beego.URLFor("UrlController.List"))
}

/**
修改url
 */
// @router /update [get,post]
func (c *UrlController) Update() {
	var url models.ServiceUrl
	filterUrls := []string{}
	id,_ := c.GetInt("id")
	if _,err := url.FindById(id);err != nil {
		c.SetSession("error","修改数据不存在")
		c.redirect(beego.URLFor("UrlController.List"))
	}
	json.Unmarshal([]byte(url.ServiceUrl),&filterUrls)
	if c.IsPost() {

		urls := c.GetStrings("service_url")
		filterUrls = filterUrls[len(filterUrls):]
		for _,v := range urls {
			if v != "" {
				filterUrls = append(filterUrls,v)
			}
		}

		if err := c.ParseForm(&url);err == nil {
			if _,err := govalidator.ValidateStruct(url);err == nil {
				if len(filterUrls) == 0 {
					c.setFlash("error", "节点不能为空")

				} else {
					serviceUrl,_ := json.Marshal(filterUrls)
					url.ServiceUrl = string(serviceUrl)
					_, err := url.Update()
					if err == nil {
						c.SetSession("success", "修改成功")
						help.Redis.Delete(beego.AppConfig.String("route::cache"))
						c.redirect(beego.URLFor("UrlController.List"))
					}
					c.setFlash("error", err.Error())
				}
			} else {
				c.setFlash("notice",err.Error())
			}
		} else {
			c.setFlash("notice",err.Error())
		}
	}

	c.Data["urls"] = filterUrls
	c.Data["urlLength"] = len(filterUrls)
	c.Data["url"] = url
	c.Data["pageTitle"] = "url修改"
	c.display()
}