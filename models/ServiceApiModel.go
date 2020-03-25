package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type ServiceApi struct {
	Id           int    `valid:"-"`
	ServiceUrlId int    `form:"service_url_id" valid:"required"`
	Method       string `form:"method" valid:"required"`
	ApiAlias     string `form:"api_alias" valid:"required"`
	ApiPath      string `form:"api_path" valid:"required"`
	InnerPath    string `form:"inner_path" valid:"required"`
	IsDelete     int8   `form:"-" valid:"-"`

	ServiceName string `form:"-" valid:"-" orm:"-"`
}

func init() {
	orm.RegisterModel(new(ServiceApi))
}

/**
根据ID查询一条
 */
func (this *ServiceApi) FindById(id int) (*ServiceApi, error) {
	err := orm.NewOrm().QueryTable(new(ServiceApi)).Filter("id", id).One(this)
	if err == nil {
		return this, nil
	}
	return nil, err
}

/**
添加serviceApi
 */
func (this *ServiceApi) Create() (int64, error) {
	return orm.NewOrm().Insert(this)
}

/**
删除serviceapi
 */
func (this *ServiceApi) Delete() (int64, error) {
	this.IsDelete = IS_DELETE_YES
	return orm.NewOrm().Update(this)
}

/**
更新serviceapi
 */
func (this *ServiceApi) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

/**
根据条件获取列表
 */
func (this *ServiceApi) ConditionList(page int, field ...interface{}) ([]ServiceApi, int64) {
	var serviceApi []ServiceApi
	pageSize, _ := beego.AppConfig.Int("pageSize")
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(new(ServiceApi))
	if len(field) > 0 {
		for i := 0; i < len(field); i = i + 2 {
			query = query.Filter(field[i].(string), field[i+1])
		}
	}
	count, _ := query.Count()
	query.Limit(pageSize).Offset(offset).OrderBy("-id").All(&serviceApi)
	return serviceApi, count
}

func (this ServiceApi) AddServiceName(apis []ServiceApi, urls []ServiceUrl) {
	for k, v := range apis {
		for _, u := range urls {
			if v.ServiceUrlId == u.Id {
				v.ServiceName = u.ServiceName
				apis[k] = v
			}
		}
	}
}

func (this ServiceApi) GetMethodAll() []string  {
	return []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
	}
}
