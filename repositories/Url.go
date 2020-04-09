package repositories

import (
	"gateway/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"gateway/service/http"
)

type urlRepository struct {
	CommonRepository
}

func NewUrlRepository() *urlRepository {
	return &urlRepository{}
}

func (c urlRepository) GetStatus() []map[string]interface{} {
	o := orm.NewOrm()
	qr := o.QueryTable(new(models.ServiceUrl))
	var urls []models.ServiceUrl
	_, err := qr.All(&urls)
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	return c.formatStatus(urls)

}

func (c urlRepository) formatStatus(data []models.ServiceUrl) []map[string]interface{} {
	ret := []map[string]interface{}{}
	for _, v := range data {
		urls := []string{}
		err := json.Unmarshal([]byte(v.ServiceUrl), &urls)
		if err != nil {
			continue
		}
		for _, item := range urls {
			tmp := map[string]interface{}{}
			tmp["name"] = v.ServiceName
			tmp["url"] = item
			if http.CheckNode(item) {
				tmp["status"] = "✅"
			} else {
				tmp["status"] = "🚫"
			}
			ret = append(ret, tmp)
		}
	}
	return ret
}
