package main

import (
	_ "myblog/routers"
	"myblog/controllers"
	"myblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

func init()  {
	models.RegisterDB()
}

func main() {
	orm.Debug = true //开启ORM调试模式
	orm.RunSyncdb("default", false, true) //自动建表

	// 注册路由
	beego.Router("/",&controllers.HomeController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/categories", &controllers.CategoryController{})

	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})   //使用beego的自动路由,如/topic？name=jim相当于/topic/jim

	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")  //添加Reply的Add方法和Delete方法
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	beego.Router("/attachment/:all", &controllers.AttachmentController{}) //文件下载
	beego.Run()
}

