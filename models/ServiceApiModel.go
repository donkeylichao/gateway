package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type ServiceApi struct {
	Id          int    `valid:"-"`
	ServiceName string `form:"service_name" valid:"required"`
	Method      string `form:"method" valid:"required"`
	ApiAlias    string `form:"api_alias" valid:"required"`
	ApiPath     string `form:"api_path" valid:"required"`
	InnerPath   string `form:"inner_path" valid:"required"`
}

/**
根据ID查询一条
 */
func (this *ServiceApi) FindById(id int) (*ServiceApi, error) {
	err := orm.NewOrm().QueryTable(TableName("service_api")).Filter("id", id).One(this)
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
	return orm.NewOrm().Delete(this)
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
func (this *ServiceApi) ConditionList(page int, field ...string) ([]*ServiceApi, int64) {
	var serviceApi []*ServiceApi
	pageSize, _ := beego.AppConfig.Int("pageSize")
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(TableName("service_api"))
	if len(field) > 0 {
		for i:=0;i<len(field) ;i=i+2  {
			query.Filter(field[i],field[i+1])
		}
	}
	count, _ := query.Count()
	query.Limit(pageSize).Offset(offset).OrderBy("-id").All(&serviceApi)
	return serviceApi, count
}
