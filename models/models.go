package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	//初始化数据库
	RegisterDB()
}

//注册数据库
func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterModelWithPrefix(beego.AppConfig.DefaultString("db::prefix", "db_"), &admin.Admin{})
	dbUser := beego.AppConfig.String("db::user")
	dbPassword := beego.AppConfig.String("db::password")
	dbDatabase := beego.AppConfig.String("db::database")
	dbCharset := beego.AppConfig.String("db::charset")
	dbHost := beego.AppConfig.String("db::host")
	dbPort := beego.AppConfig.String("db::port")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase, dbCharset)
	maxIdle := beego.AppConfig.DefaultInt("db::maxIdle", 50)
	maxConn := beego.AppConfig.DefaultInt("db::maxConn", 300)
	if err := orm.RegisterDataBase("default", "mysql", dbLink, maxIdle, maxConn); err != nil {
		panic(err)
	}

}
