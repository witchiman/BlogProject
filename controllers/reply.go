package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add()  {
	if !checkLogin(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	topicId := c.Input().Get("topicId")
	nickName := c.Input().Get("nickName")
	content := c.Input().Get("content")

	err := models.AddReply(topicId, nickName, content)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic/view/"+topicId, 302)

}

func (c *ReplyController) Delete()  {
	topicId := c.Input().Get("topicId")
	replyId := c.Input().Get("replyId")

	err := models.DeleteReply(replyId)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic/view/"+topicId, 302)
}
