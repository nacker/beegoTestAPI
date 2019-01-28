package main

import (
	_ "beegoTestAPI/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beegoTestAPI/models"
)

func init()  {
	// 初始化数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beegoTestAPI?charset=utf8")

	// 注册model
	orm.RegisterModel(new(models.User))

	// create table
	//第二个参数是强制更新数据库
	//第三个参数是如果没有则同步
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
