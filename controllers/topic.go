package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	"fmt"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get()  {
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	c.Data["IsLogin"] = checkLogin(c.Ctx)

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
}

func (c *TopicController) Post()  {
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	titleId := c.Input().Get("titleId")
	fmt.Println("titleId", titleId)

	var err error
	if len(titleId)!=0 {  //如是参数里含有文章id，说明操作是修改文章
		err = models.ModifyTopic(titleId, title, content)
	}else {
		err = models.AddTopic(title, content, category)
	}

	if err != nil {
		beego.Error(err.Error())
	}
	c.Redirect("/", 302)
}


func(c *TopicController) Add()  {
	c.TplName = "topic_add.html"
}

func (c *TopicController) Delete()  {
	id := c.Ctx.Input.Param("0")  //通过自动路由获取值,如/topic/delete/2,相当于/topic?op=delete&&value=2
	err := models.DeleteTopic(id)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)

}

func (c *TopicController) Modify()  {
	titleId := c.Ctx.Input.Param("0")
	c.Data["TitleId"] = titleId

	topic, err := models.GetTopic(titleId)
	if err!=nil {
		beego.Error(err)
	}

	c.Data["Topic"] = topic
	c.TplName = "topic_modify.html"

}

func (c *TopicController) View()  {
	c.Data["IsTopic"] = true
	c.TplName = "topic_view.html"

	titleId := c.Ctx.Input.Param("0")
	fmt.Println("View",titleId)
	topic, err := models.GetTopic(titleId)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic

}
