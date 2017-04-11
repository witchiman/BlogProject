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
	topics,err := models.GetAllTopics(category, true)
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
