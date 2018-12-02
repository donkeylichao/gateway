package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type User struct {
	Id            int    `valid:"-"`
	Name          string `form:"name" valid:"required"`
	Email         string `form:"email" valid:"email,required"`
	Password      string `valid:"-"`
	Phone         string `form:"phone" valid:"-"`
	RealName      string `form:"real_name" valid:"-"`
	LastLoginIp   string `valid:"ip"`
	LastLoginTime int    `valid:"-"`
}

//通过邮箱密码获取用户信息
func (this *User) GetUserByEmail(user ...string) (*User, error) {

	o := orm.NewOrm()
	err := o.QueryTable(new(User)).Filter("email", user[0]).Filter("password", user[1]).RelatedSel().One(this)
	if err == nil {
		return this, nil
	}
	return nil, err
}

/**
根据ID获取用户信息
 */
func (this *User) FindById(id int) (*User, error) {
	o := orm.NewOrm()
	err := o.QueryTable(new(User)).Filter("id", id).One(this)
	//panic(err)
	if err == nil {
		return this, nil
	}
	return nil, err
}

/**
条件获取列表
 */
func (this *User) ConditionList(page int, fields []string) ([]*User, int64) {
	var users []*User
	pageSize, _ := beego.AppConfig.Int("pageSize")
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(TableName("user"))
	count, _ := query.Count()
	query.Limit(pageSize).Offset(offset).OrderBy("-id").All(&users)
	return users, count
}

/**
添加用户方法
 */
func (this *User) Create() (int64, error) {
	return orm.NewOrm().Insert(this)
}

/**
删除
 */
func (this *User) Del() (int64, error) {
	return orm.NewOrm().Delete(this)
}

/**
通过ID删除
 */

func (this *User) Delete(id int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).Delete()
}

/**
修改
 */
func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}
