package main

import (
	_ "myblog/routers"
	"github.com/astaxie/beego"
	"myblog/models"
	"github.com/astaxie/beego/orm"
	"myblog/controllers"
)

func init()  {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Router("/",&controllers.MainController{})
	beego.Run()
}

