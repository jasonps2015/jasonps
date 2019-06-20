package controllers


func (c *LayuiadminGOController) UserLogin() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/user/login.html"
}

func (c *LayuiadminGOController) UserForget() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.Data["search_type"] = "resetpass"
	c.Data["search_type2"] = "resetpass"

	c.TplName = "layuiadmingo/user/forget.html"
}

func (c *LayuiadminGOController) UserReg() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/user/reg.html"
}

func (c *LayuiadminGOController) UserUserList() {
	c.Data["d_avatar"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/user/user/list.html"
}

func (c *LayuiadminGOController) UserUserUserform() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/user/user/userform.html"
}

func (c *LayuiadminGOController) UserAdministratorsAdminform() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/user/administrators/adminform.html"
}

func (c *LayuiadminGOController) UserAdministratorsList() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.Data["d_check"] = true
	//c.Data["d_role"] = "超级管理员"
	c.Data["if_d_role"] = "超级管理员"
	c.Data["d_role"] = true

	c.TplName = "layuiadmingo/user/administrators/list.html"
}

func (c *LayuiadminGOController)UserAdministratorsRole() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["d_check"] = true
	c.TplName = "layuiadmingo/user/administrators/role.html"
}

func (c *LayuiadminGOController) UserAdministratorsRoleform() {
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/user/administrators/roleform.html"
}


