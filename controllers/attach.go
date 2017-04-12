package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"io"
	"net/url"
)

type AttachmentController struct {
	beego.Controller
}

func (c *AttachmentController) Get()  {
	//获取文件路径
	path, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:])
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	file, err := os.Open(path)
	if err!=nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	defer file.Close()

	_, err = io.Copy(c.Ctx.ResponseWriter, file)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}
