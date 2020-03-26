package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type ServiceUrl struct {
	Id          int           `valid:"-" json:"id"`
	ServiceName string        `form:"service_name" valid:"required" json:"service_name"`
	ServiceUrl  string        `form:"service_url" valid:"required"`
}

func init() {
	orm.RegisterModel(new(ServiceUrl))
}

/**
根据ID获取一条
 */
func (this *ServiceUrl) FindById(id int) (*ServiceUrl, error) {
	err := orm.NewOrm().QueryTable(new(ServiceUrl)).Filter("id", id).One(this)
	if err == nil {
		return this, nil
	}
	return nil, err
}

/**
获取列表
 */
func (this *ServiceUrl) ConditionList(page int, field ...interface{}) ([]*ServiceUrl, int64) {
	var serviceUrl []*ServiceUrl
	pageSize, _ := beego.AppConfig.Int("pageSize")
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(new(ServiceUrl))
	if len(field) > 0 {
		for i := 0; i < len(field); i = i + 2 {
			query = query.Filter(field[i].(string), field[i+1])
		}
	}
	count, _ := query.Count()
	query.Limit(pageSize).Offset(offset).OrderBy("-id").All(&serviceUrl)

	return serviceUrl, count
}

/**
添加
 */
func (this *ServiceUrl) Create() (int64, error) {
	return orm.NewOrm().Insert(this)
}

/**
删除
 */
func (this *ServiceUrl) Delete() (int64, error) {
	return orm.NewOrm().Delete(this)
}

/**
更新
 */
func (this *ServiceUrl) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

/**
获取所有列表
 */
func (*ServiceUrl) IdAndUrlList() map[int]string {
	var serviceUrl []*ServiceUrl
	orm.NewOrm().QueryTable(new(ServiceUrl)).All(&serviceUrl)

	data := make(map[int]string)
	for _, u := range serviceUrl {
		data[u.Id] = u.ServiceUrl
	}
	return data
}