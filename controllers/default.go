package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.TplName = "home.html"
	c.Data["IsHome"] = true
	topics,err := models.GetAllTopics(true)
	if err!=nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
	c.Data["IsLogin"] = checkLogin(c.Ctx)

}
