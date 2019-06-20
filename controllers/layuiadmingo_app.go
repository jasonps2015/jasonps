package controllers

//App Content
func (c *LayuiadminGOController) AppContentComment() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/comment.html"
}


func (c *LayuiadminGOController) AppContentContform() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/contform.html"
}


func (c *LayuiadminGOController) AppContentList() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["d_status"] = true

	c.TplName = "layuiadmingo/app/content/list.html"
}


func (c *LayuiadminGOController) AppContentListform() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/listform.html"
}


func (c *LayuiadminGOController) AppContentTags() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/tags.html"
}


func (c *LayuiadminGOController) AppContentTagsform() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/tagsform.html"
}

//App Forum
func (c *LayuiadminGOController) AppForumList() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["d_avatar"] = "layuiAdmin-v1.2.1 std"
	c.Data["d_top"] = true

	c.TplName = "layuiadmingo/app/forum/list.html"
}


func (c *LayuiadminGOController) AppForumListform() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/forum/listform.html"
}


func (c *LayuiadminGOController) AppForumReplys() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["d_avatar"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/forum/replys.html"
}


func (c *LayuiadminGOController) AppForumReplysform() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/forum/replysform.html"
}

//App  Mall
func (c *LayuiadminGOController) AppMallCategory() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/mall/category.html"
}


func (c *LayuiadminGOController) AppMallList() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/mall/list.html"
}

func (c *LayuiadminGOController) AppMallSpecs() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/mall/specs.html"
}







//message
func (c *LayuiadminGOController) AppMessageDetail() {
	c.Data["d_data_title"] = "修改名字"
	c.Data["d_data_time"] = "20190621 03:34:05"
	c.Data["d_data_content"] = "修改名字 layuiAdmin-v1.2.1 std"
	c.Data["d_data_day"] = 20190621

	c.TplName = "layuiadmingo/app/content/comment.html"
}


func (c *LayuiadminGOController) AppMessageIndex() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/comment.html"
}

//workorder
func (c *LayuiadminGOController) AppWorkorderList() {
	c.Data["d_state"] = "已处理"
	c.Data["d_stateyes"] = "已处理"
	c.Data["d_stateno"] = "未分配"
	c.Data["d_orderid"] = "2.4.5"
	c.Data["d_progress"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/comment.html"
}


func (c *LayuiadminGOController) AppWorkorderListform() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["layui_admin_v"] = "layuiAdmin-v1.2.1 std"
	c.Data["layui_v"] = "2.4.5"

	c.TplName = "layuiadmingo/app/content/comment.html"
}



