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

	category := c.Input().Get("category")
	label := c.Input().Get("label")
	topics,err := models.GetAllTopics(category, label, true)
	if err!=nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics

	categories, err := models.GetCategories()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categories"] = categories
	c.Data["IsLogin"] = checkLogin(c.Ctx)

}
