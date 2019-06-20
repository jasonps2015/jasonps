package controllers


func (c *LayuiadminGOController) SeniorEchartsBar() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/senior/echarts/bar.html"
}

func (c *LayuiadminGOController) SeniorEchartsLine() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/senior/echarts/line.html"
}

func (c *LayuiadminGOController) SeniorEchartsMap() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/senior/echarts/map.html"
}

func (c *LayuiadminGOController) SeniorImIndex() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/senior/im/index.html"
}