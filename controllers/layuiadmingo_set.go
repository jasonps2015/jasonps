package controllers


func (c *LayuiadminGOController) SetUserInfo() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/set/user/info.html"
}

func (c *LayuiadminGOController) SetUserPassword() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/set/user/password.html"
}

func (c *LayuiadminGOController) SetSystemEmail() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/set/system/email.html"
}

func (c *LayuiadminGOController) SetSystemWebsite() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/set/system/website.html"
}
