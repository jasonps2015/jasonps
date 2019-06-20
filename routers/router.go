package routers

import (
	"github.com/astaxie/beego"
	"jasonps/controllers"
	layuiadmingo "jasonps/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/admingo", &admingo.AdminGOController{})
	//beego.Router("/layuiadmingo", &layuiadmingo.LayuiadminGOController{})

	beego.Router("/layuiadmingo", &layuiadmingo.LayuiadminGOController{},"*:Get")

	//home
	beego.Router("/home/console", &layuiadmingo.LayuiadminGOController{},"*:HomeConsole")
	beego.Router("/home/homepage1", &layuiadmingo.LayuiadminGOController{},"*:HomeHomepage1")
	beego.Router("/home/homepage2", &layuiadmingo.LayuiadminGOController{},"*:HomeHomepage2")


  //user
	beego.Router("/user", &layuiadmingo.LayuiadminGOController{},"*:UserLogin")
	beego.Router("/user/login", &layuiadmingo.LayuiadminGOController{},"*:UserLogin")
	beego.Router("/user/forget", &layuiadmingo.LayuiadminGOController{},"*:UserForget")
	beego.Router("/user/reg", &layuiadmingo.LayuiadminGOController{},"*:UserReg")
	beego.Router("/user/user/list", &layuiadmingo.LayuiadminGOController{},"*:UserUserList")
	beego.Router("/user/user/form", &layuiadmingo.LayuiadminGOController{},"*:UserUserUserform")
	beego.Router("/user/administrators/adminform", &layuiadmingo.LayuiadminGOController{},"*:UserAdministratorsAdminform")
	beego.Router("/user/administrators/list", &layuiadmingo.LayuiadminGOController{},"*:UserAdministratorsList")
	beego.Router("/user/administrators/role", &layuiadmingo.LayuiadminGOController{},"*:UserAdministratorsRole")
	beego.Router("/user/administrators/roleform", &layuiadmingo.LayuiadminGOController{},"*:UserAdministratorsRoleform")

	//set system
	beego.Router("/set", &layuiadmingo.LayuiadminGOController{},"*:SetUserInfo")
	beego.Router("/set/user/info", &layuiadmingo.LayuiadminGOController{},"*:SetUserInfo")
	beego.Router("/set/user/password", &layuiadmingo.LayuiadminGOController{},"*:SetUserPassword")
	beego.Router("/set/system/email", &layuiadmingo.LayuiadminGOController{},"*:SetSystemEmail")
	beego.Router("/set/system/website", &layuiadmingo.LayuiadminGOController{},"*:SetSystemWebsite")


	//Senior
	beego.Router("/senior", &layuiadmingo.LayuiadminGOController{},"*:SeniorEchartsBar")
	beego.Router("/senior/echarts/bar", &layuiadmingo.LayuiadminGOController{},"*:SeniorEchartsBar")
	beego.Router("/senior/echarts/line", &layuiadmingo.LayuiadminGOController{},"*:SeniorEchartsLine")
	beego.Router("/senior/echarts/map", &layuiadmingo.LayuiadminGOController{},"*:SeniorEchartsMap")
	beego.Router("/senior/im/index", &layuiadmingo.LayuiadminGOController{},"*:SeniorImIndex")



	//app  content
	beego.Router("/app", &layuiadmingo.LayuiadminGOController{},"*:AppContentComment")
	beego.Router("/app/content/comment", &layuiadmingo.LayuiadminGOController{},"*:AppContentComment")
	beego.Router("/app/content/contform", &layuiadmingo.LayuiadminGOController{},"*:AppContentContform")
	beego.Router("/app/content/list", &layuiadmingo.LayuiadminGOController{},"*:AppContentList")
	beego.Router("/app/content/listform", &layuiadmingo.LayuiadminGOController{},"*:AppContentListform")
	beego.Router("/app/content/tags", &layuiadmingo.LayuiadminGOController{},"*:AppContentTags")
	beego.Router("/app/content/tagsform", &layuiadmingo.LayuiadminGOController{},"*:AppContentTagsform")
	//app  forum
	beego.Router("/app/forum/list", &layuiadmingo.LayuiadminGOController{},"*:AppForumList")
	beego.Router("/app/forum/listform", &layuiadmingo.LayuiadminGOController{},"*:AppForumListform")
	beego.Router("/app/forum/replys", &layuiadmingo.LayuiadminGOController{},"*:AppForumReplys")
	beego.Router("/app/forum/replysform", &layuiadmingo.LayuiadminGOController{},"*:AppForumReplysform")
	//app  mail
	beego.Router("/app/mail/category", &layuiadmingo.LayuiadminGOController{},"*:AppMallCategory")
	beego.Router("/app/mail/list", &layuiadmingo.LayuiadminGOController{},"*:AppMallList")
	beego.Router("/app/mail/specs", &layuiadmingo.LayuiadminGOController{},"*:AppMallSpecs")
	//app  message
	beego.Router("/app/message/detail", &layuiadmingo.LayuiadminGOController{},"*:AppMessageDetail")
	beego.Router("/app/message/index", &layuiadmingo.LayuiadminGOController{},"*:AppMessageIndex")
	//app  workorder
	beego.Router("/app/workorder/list", &layuiadmingo.LayuiadminGOController{},"*:AppWorkorderList")
	beego.Router("/app/workorder/listform", &layuiadmingo.LayuiadminGOController{},"*:AppWorkorderListform")



	//template
	beego.Router("/template", &layuiadmingo.LayuiadminGOController{},"*:TemplateAddresslist")
	beego.Router("/template/addresslist", &layuiadmingo.LayuiadminGOController{},"*:TemplateAddresslist")
	beego.Router("/template/caller", &layuiadmingo.LayuiadminGOController{},"*:TemplateCaller")
	beego.Router("/template/goodslist", &layuiadmingo.LayuiadminGOController{},"*:TemplateGoodslist")
	beego.Router("/template/msgboard", &layuiadmingo.LayuiadminGOController{},"*:TemplateMsgboard")
	beego.Router("/template/personalpage", &layuiadmingo.LayuiadminGOController{},"*:TemplatePersonalpage")
	beego.Router("/template/search", &layuiadmingo.LayuiadminGOController{},"*:TemplateSearch")
	beego.Router("/template/temp", &layuiadmingo.LayuiadminGOController{},"*:TemplateTemp")
	//template tips
	beego.Router("/template/tips/404", &layuiadmingo.LayuiadminGOController{},"*:TemplateTips404")
	beego.Router("/template/tips/error", &layuiadmingo.LayuiadminGOController{},"*:TemplateTempError")

	//
	//beego.Router("/template", &layuiadmingo.LayuiadminGOController{},"*:Template")
	//beego.Router("/template", &layuiadmingo.LayuiadminGOController{},"*:Template")
	//beego.Router("/template", &layuiadmingo.LayuiadminGOController{},"*:Template")
	//beego.Router("/template", &layuiadmingo.LayuiadminGOController{},"*:Template")

}
