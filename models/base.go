package models

import (
	"crypto/md5"
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
	"strings"
	"time"
)

func init() {
	//dbhost := beego.AppConfig.String("dbhost")
	//dbport := beego.AppConfig.String("dbport")
	//dbuser := beego.AppConfig.String("dbuser")
	//dbpassword := beego.AppConfig.String("dbpassword")
	//dbname := beego.AppConfig.String("dbname")
	//if dbport == "" {
	//	dbport = "3306"
	//}
	//dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDataBase("default", "sqlite3", "nginx.db")
	//orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(Domain), new(Option))
	//, new(Option))
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}
