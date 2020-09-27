package main

import (
	//_"BeegoProject0603(1)/BeegoProject0603/db_mysql"
	_ "Project922/routers"
	_ "database/sql"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
)

func main()  {
	beego.Run()
}
//var Db *sql.DB

	//appName :=config.String("appname")
	//fmt.Println("应用名称:",appName)
	//port,err :=config.Int("httpport")
	//if err!=nil {
	//	panic("项目配置文件解析失败，请检查配置文件")
	//}
	//dbDirver :=config.String("db_driverName")
	//dbUser:=config.String("db_user")
	//dbPassword:=config.String("db_password")
	//dbIp:=config.String("db_ip")
	////dbPort:=config.String("db_port")
	//dbName:=config.String("db_name")
	//
	//connUrl :=dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"charset=uft-8"
	//fmt.Println(port)
	//  db,err:=sql.Open(dbDirver,connUrl)
	//
	//beego.Run()


//在初始化函数中链接数据库


//leetcode


