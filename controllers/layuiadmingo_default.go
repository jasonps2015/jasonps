package controllers

import (
	"github.com/astaxie/beego"
)

type LayuiadminGOController struct {
	beego.Controller
}

func (c *LayuiadminGOController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "layuiadmingo/index.html"
}
