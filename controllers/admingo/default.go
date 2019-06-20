package controllers

import (
	"github.com/astaxie/beego"
)

type AdminGOController struct {
	beego.Controller
}

func (c *AdminGOController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
