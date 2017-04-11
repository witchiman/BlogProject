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

	topics, err := models.GetAllTopics("", false)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
}

func (c *TopicController) Post()  {
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	category := c.Input().Get("category")
	topicId := c.Input().Get("topicId")
	fmt.Println("topicId", topicId)

	var err error
	if len(topicId)!=0 {  //如是参数里含有文章id，说明操作是修改文章
		err = models.ModifyTopic(topicId, title, content, category)
		c.Redirect("/topic", 302)
	}else {
		err = models.AddTopic(title, content, category)
		c.Redirect("/topic", 302)

	}

	if err != nil {
		beego.Error(err.Error())
	}
}


func(c *TopicController) Add()  {
	if !checkLogin(c.Ctx) {       //检查是否需要重新登录
		c.Redirect("/login", 302)
		return
	}

	c.TplName = "topic_add.html"
}

func (c *TopicController) Delete()  {
	if !checkLogin(c.Ctx) {       //检查是否需要重新登录
		c.Redirect("/login", 302)
		return
	}

	id := c.Ctx.Input.Param("0")  //通过自动路由获取值,如/topic/delete/2,相当于/topic?op=delete&&value=2
	err := models.DeleteTopic(id)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)

}

func (c *TopicController) Modify()  {
	if !checkLogin(c.Ctx) {       //检查是否需要重新登录
		c.Redirect("/login", 302)
		return
	}

	topicId := c.Ctx.Input.Param("0")
	c.Data["TopicId"] = topicId

	topic, err := models.GetTopic(topicId)
	if err!=nil {
		beego.Error(err)
	}

	c.Data["Topic"] = topic
	c.TplName = "topic_modify.html"

}

func (c *TopicController) View()  {
	c.Data["IsTopic"] = true
	c.TplName = "topic_view.html"
	c.Data["IsLogin"] = checkLogin(c.Ctx)

	topicId := c.Ctx.Input.Param("0")
	fmt.Println("View", topicId)
	topic, err := models.GetTopic(topicId)

	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Topic"] = topic

	replies, err := models.GetAllReplies(topicId)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Replies"] = replies

}
