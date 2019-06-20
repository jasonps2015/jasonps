package controllers


func (c *LayuiadminGOController) HomeConsole() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/home/console.html"
}

func (c *LayuiadminGOController) HomeHomepage1() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/home/homepage1.html"
}

func (c *LayuiadminGOController) HomeHomepage2() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/home/homepage2.html"
}
