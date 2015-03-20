package models

import (
	"github.com/astaxie/beego/orm"
)

type Option struct {
	Id    int64  `orm:"column(id);auto"`
	Name  string `orm:"column(name)"`
	Value string `orm:"column(value)"`
}

//高级查询
func (m *Option) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

//读取字段相同的记录
func (m *Option) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func AddOption(m *Option) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
