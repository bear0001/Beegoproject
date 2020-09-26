package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)
var Db*sql.DB

func init()  {
	fmt.Println("连接mysql数据库")
	config :=beego.AppConfig
	dbDriver :=config.String("db_riverName")
	dbUser :=config.String("db_user")
	dbPassword:=config.String("db_password")
	dbIp:=config.String("db_ip")
	dbName :=config.String("db_name")

	connUrl:=dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db,err:=sql.Open(dbDriver,connUrl)
	if err !=nil{
		panic("数据库连接错误，请配置检查")
    Db = db
}


