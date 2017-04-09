package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type CategoryController struct{
	beego.Controller
}

func (c *CategoryController) Get() {
	option := c.Input().Get("op")
	switch option {
	case "add":
		name := c.Input().Get("category")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/categories", 302)

	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err!=nil {
			beego.Error(err)
		}
		c.Redirect("/categories", 302)

	}

	categories, err := models.GetCategories()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categories"] = categories
	c.Data["IsLogin"] = checkLogin(c.Ctx)
	c.Data["IsCategory"] = true
	c.TplName = "category.html"
}

