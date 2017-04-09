package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController)Get() {
	if c.Input().Get("exit") == "true"{
		c.Ctx.SetCookie("uname", "", -1, "/" )
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 302)
	}

	c.TplName = "login.html"
}

func(c *LoginController) Post() {
	name := c.Input().Get("uname")
	password := c.Input().Get("pwd")
	autologin := c.Input().Get("autologin") == "on"
	fmt.Println("the name is ", name)

	if (name==beego.AppConfig.String("adminName")) &&   // 是否与配置文件相同
		(password==beego.AppConfig.String("adminPassword")){
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}

		c.Ctx.SetCookie("uname", name, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}

	c.Redirect("/", 302)
	return
}

func  checkLogin(ctx *context.Context) bool {
	cookie, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	name := cookie.Value

	cookie, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := cookie.Value

	result := (name==beego.AppConfig.String("adminName"))&&
		(password==beego.AppConfig.String("adminPassword"))
	fmt.Println("The check result is ", result)

	return result
}
