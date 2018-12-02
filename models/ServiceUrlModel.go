package models

type serviceUrl struct {
	Id          int    `valid:"-"`
	ServiceName string `form:"service_name" valid:"required"`
	ServiceUrl  string `form:"service_url" valid:"required"`
}
