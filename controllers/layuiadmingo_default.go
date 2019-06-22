package controllers

import (
	"github.com/astaxie/beego"
)

type LayuiadminGOController struct {
	beego.Controller
}

type Tree struct {
	Id         int64      `json:"id"`
	Text       string     `json:"text"`
	IconCls    string     `json:"iconCls"`
	Checked    string     `json:"checked"`
	State      string     `json:"state"`
	Children   []Tree     `json:"children"`
	Attributes Attributes `json:"attributes"`
}
type Attributes struct {
	Url   string `json:"url"`
	Price int64  `json:"price"`
}



func (c *LayuiadminGOController) Get() {

	userinfo :=c.GetSession("userinfo")
	if userinfo ==nil{
		//获得默认登录地址，如用户没有登录进行  302 重置URL
		c.Ctx.Redirect(302,beego.AppConfig.String("rbac_auth_gateway"))
	}
	c.Data["Website"] = "beego.me"

	c.TplName = "layuiadmingo/index.html"
}

