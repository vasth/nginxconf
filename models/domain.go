package models

import (
	"github.com/astaxie/beego/orm"
)

type Domain struct {
	Id          int64  `orm:"column(id);auto"`
	Server_name string `orm:"column(server_name)"`
	Port        int64  `orm:"column(port)"`
	Proxy_pass  string `orm:"column(proxy_pass)"`
	Access_log  string `orm:"column(access_log)"`
	Expires     string `orm:"column(expires)"`
	Root        string `orm:"column(root)"`
}

//高级查询
func (m *Domain) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

//读取字段相同的记录
func (m *Domain) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func AddDomain(m *Domain) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
