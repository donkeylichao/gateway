package models

type serviceApi struct {
	Id          int    `valid:"-"`
	ServiceName string `form:"service_name" valid:"required"`
	Method      string `form:"method" valid:"required"`
	ApiAlias    string `form:"api_alias" valid:"required"`
	ApiPath     string `form:"api_path" valid:"required"`
	InnerPath   string `form:"inner_path" valid:"required"`
}
