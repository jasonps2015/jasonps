// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"jasonps/controllers"

	"github.com/astaxie/beego"
)

func init() {


	beego.Router("/layuiadmingo", &controllers.LayuiadminGOController{},"*:Get")



	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/group",			beego.NSInclude(&controllers.GroupController{},		),),
		beego.NSNamespace("/node",			beego.NSInclude(&controllers.NodeController{},		),),
		beego.NSNamespace("/node_roles",	beego.NSInclude(&controllers.NodeRolesController{},		),),
		beego.NSNamespace("/role",			beego.NSInclude(&controllers.RoleController{},		),),
		beego.NSNamespace("/user",			beego.NSInclude(&controllers.UserController{},		),),
		beego.NSNamespace("/user_roles",	beego.NSInclude(&controllers.UserRolesController{},		),),
		//beego.nsn
	)




	beego.AddNamespace(ns)




	//home
	beego.Router("/home/console", &controllers.LayuiadminGOController{},"*:HomeConsole")
	beego.Router("/home/homepage1", &controllers.LayuiadminGOController{},"*:HomeHomepage1")
	beego.Router("/home/homepage2", &controllers.LayuiadminGOController{},"*:HomeHomepage2")


	//user
	beego.Router("/user", &controllers.LayuiadminGOController{},"*:UserLogin")
	beego.Router("/user/login", &controllers.LayuiadminGOController{},"*:UserLogin")
	beego.Router("/user/forget", &controllers.LayuiadminGOController{},"*:UserForget")
	beego.Router("/user/reg", &controllers.LayuiadminGOController{},"*:UserReg")
	beego.Router("/user/user/list", &controllers.LayuiadminGOController{},"*:UserUserList")
	beego.Router("/user/user/form", &controllers.LayuiadminGOController{},"*:UserUserUserform")
	beego.Router("/user/administrators/adminform", &controllers.LayuiadminGOController{},"*:UserAdministratorsAdminform")
	beego.Router("/user/administrators/list", &controllers.LayuiadminGOController{},"*:UserAdministratorsList")
	beego.Router("/user/administrators/role", &controllers.LayuiadminGOController{},"*:UserAdministratorsRole")
	beego.Router("/user/administrators/roleform", &controllers.LayuiadminGOController{},"*:UserAdministratorsRoleform")

	//set system
	beego.Router("/set", &controllers.LayuiadminGOController{},"*:SetUserInfo")
	beego.Router("/set/user/info", &controllers.LayuiadminGOController{},"*:SetUserInfo")
	beego.Router("/set/user/password", &controllers.LayuiadminGOController{},"*:SetUserPassword")
	beego.Router("/set/system/email", &controllers.LayuiadminGOController{},"*:SetSystemEmail")
	beego.Router("/set/system/website", &controllers.LayuiadminGOController{},"*:SetSystemWebsite")


	//Senior
	beego.Router("/senior", &controllers.LayuiadminGOController{},"*:SeniorEchartsBar")
	beego.Router("/senior/echarts/bar", &controllers.LayuiadminGOController{},"*:SeniorEchartsBar")
	beego.Router("/senior/echarts/line", &controllers.LayuiadminGOController{},"*:SeniorEchartsLine")
	beego.Router("/senior/echarts/map", &controllers.LayuiadminGOController{},"*:SeniorEchartsMap")
	beego.Router("/senior/im/index", &controllers.LayuiadminGOController{},"*:SeniorImIndex")



	//app  content
	beego.Router("/app", &controllers.LayuiadminGOController{},"*:AppContentComment")
	beego.Router("/app/content/comment", &controllers.LayuiadminGOController{},"*:AppContentComment")
	beego.Router("/app/content/contform", &controllers.LayuiadminGOController{},"*:AppContentContform")
	beego.Router("/app/content/list", &controllers.LayuiadminGOController{},"*:AppContentList")
	beego.Router("/app/content/listform", &controllers.LayuiadminGOController{},"*:AppContentListform")
	beego.Router("/app/content/tags", &controllers.LayuiadminGOController{},"*:AppContentTags")
	beego.Router("/app/content/tagsform", &controllers.LayuiadminGOController{},"*:AppContentTagsform")
	//app  forum
	beego.Router("/app/forum/list", &controllers.LayuiadminGOController{},"*:AppForumList")
	beego.Router("/app/forum/listform", &controllers.LayuiadminGOController{},"*:AppForumListform")
	beego.Router("/app/forum/replys", &controllers.LayuiadminGOController{},"*:AppForumReplys")
	beego.Router("/app/forum/replysform", &controllers.LayuiadminGOController{},"*:AppForumReplysform")
	//app  mail
	beego.Router("/app/mail/category", &controllers.LayuiadminGOController{},"*:AppMallCategory")
	beego.Router("/app/mail/list", &controllers.LayuiadminGOController{},"*:AppMallList")
	beego.Router("/app/mail/specs", &controllers.LayuiadminGOController{},"*:AppMallSpecs")
	//app  message
	beego.Router("/app/message/detail", &controllers.LayuiadminGOController{},"*:AppMessageDetail")
	beego.Router("/app/message/index", &controllers.LayuiadminGOController{},"*:AppMessageIndex")
	//app  workorder
	beego.Router("/app/workorder/list", &controllers.LayuiadminGOController{},"*:AppWorkorderList")
	beego.Router("/app/workorder/listform", &controllers.LayuiadminGOController{},"*:AppWorkorderListform")



	//template
	beego.Router("/template", &controllers.LayuiadminGOController{},"*:TemplateAddresslist")
	beego.Router("/template/addresslist", &controllers.LayuiadminGOController{},"*:TemplateAddresslist")
	beego.Router("/template/caller", &controllers.LayuiadminGOController{},"*:TemplateCaller")
	beego.Router("/template/goodslist", &controllers.LayuiadminGOController{},"*:TemplateGoodslist")
	beego.Router("/template/msgboard", &controllers.LayuiadminGOController{},"*:TemplateMsgboard")
	beego.Router("/template/personalpage", &controllers.LayuiadminGOController{},"*:TemplatePersonalpage")
	beego.Router("/template/search", &controllers.LayuiadminGOController{},"*:TemplateSearch")
	beego.Router("/template/temp", &controllers.LayuiadminGOController{},"*:TemplateTemp")
	//template tips
	beego.Router("/template/tips/404", &controllers.LayuiadminGOController{},"*:TemplateTips404")
	beego.Router("/template/tips/error", &controllers.LayuiadminGOController{},"*:TemplateTempError")

	//iframe
	beego.Router("/iframe", &controllers.LayuiadminGOController{},"*:IframeLayerIframe")
	beego.Router("/iframe/layer/iframe", &controllers.LayuiadminGOController{},"*:IframeLayerIframe")


	//system
	beego.Router("/system", &controllers.LayuiadminGOController{},"*:SystemAbout")
	beego.Router("/system/about", &controllers.LayuiadminGOController{},"*:SystemAbout")
	beego.Router("/system/get", &controllers.LayuiadminGOController{},"*:SystemGet")
	beego.Router("/system/more", &controllers.LayuiadminGOController{},"*:SystemMore")
	beego.Router("/system/theme", &controllers.LayuiadminGOController{},"*:SystemTheme")

	//component
	beego.Router("/component/", &controllers.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/anim", &controllers.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/auxiliar", &controllers.LayuiadminGOController{},"*:ComponentAuxiliarIndex")
	beego.Router("/component/badge", &controllers.LayuiadminGOController{},"*:ComponentBadgeIndex")
	beego.Router("/component/button", &controllers.LayuiadminGOController{},"*:ComponentButtonIndex")
	beego.Router("/component/carousel", &controllers.LayuiadminGOController{},"*:ComponentCarouselIndex")
	beego.Router("/component/code", &controllers.LayuiadminGOController{},"*:ComponentCodeIndex")
	beego.Router("/component/colorpicker", &controllers.LayuiadminGOController{},"*:ComponentColorpickerIndex")
	beego.Router("/component/flow", &controllers.LayuiadminGOController{},"*:ComponentFlowIndex")

	beego.Router("/component/index", &controllers.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/anim/index", &controllers.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/auxiliar/index", &controllers.LayuiadminGOController{},"*:ComponentAuxiliarIndex")
	beego.Router("/component/badge/index", &controllers.LayuiadminGOController{},"*:ComponentBadgeIndex")
	beego.Router("/component/button/index", &controllers.LayuiadminGOController{},"*:ComponentButtonIndex")
	beego.Router("/component/carousel/index", &controllers.LayuiadminGOController{},"*:ComponentCarouselIndex")
	beego.Router("/component/code/index", &controllers.LayuiadminGOController{},"*:ComponentCodeIndex")
	beego.Router("/component/colorpicker/index", &controllers.LayuiadminGOController{},"*:ComponentColorpickerIndex")
	beego.Router("/component/flow/index", &controllers.LayuiadminGOController{},"*:ComponentFlowIndex")

	//component form
	beego.Router("/component/form", &controllers.LayuiadminGOController{},"*:ComponentFormElement")
	beego.Router("/component/form/element", &controllers.LayuiadminGOController{},"*:ComponentFormElement")
	beego.Router("/component/form/group", &controllers.LayuiadminGOController{},"*:ComponentFormGroup")
	//component grid
	beego.Router("/component/grid", &controllers.LayuiadminGOController{},"*:ComponentGridAll")
	beego.Router("/component/grid/all", &controllers.LayuiadminGOController{},"*:ComponentGridAll")
	beego.Router("/component/grid/list", &controllers.LayuiadminGOController{},"*:ComponentGridList")
	beego.Router("/component/grid/mobile", &controllers.LayuiadminGOController{},"*:ComponentGridMobile")
	beego.Router("/component/grid/mobile-pc", &controllers.LayuiadminGOController{},"*:ComponentGridMobile_pc")
	beego.Router("/component/grid/speed-dial", &controllers.LayuiadminGOController{},"*:ComponentGridSpeed_dial")
	beego.Router("/component/grid/stack", &controllers.LayuiadminGOController{},"*:ComponentGridStack")



	//component laydate
	beego.Router("/component/laydate", &controllers.LayuiadminGOController{},"*:ComponentLaydateIndex")
	beego.Router("/component/laydate/index", &controllers.LayuiadminGOController{},"*:ComponentLaydateIndex")
	beego.Router("/component/laydate/demo1", &controllers.LayuiadminGOController{},"*:ComponentLaydateDemo1")
	beego.Router("/component/laydate/demo2", &controllers.LayuiadminGOController{},"*:ComponentLaydateDemo2")
	beego.Router("/component/laydate/special-demo", &controllers.LayuiadminGOController{},"*:ComponentLaydateSpecial_demo")
	beego.Router("/component/laydate/theme", &controllers.LayuiadminGOController{},"*:ComponentLaydateTheme")

	//beego.Router("/component/layer", &layuiadmingo.LayuiadminGOController{},"*:ComponentLayerList")
	beego.Router("/component/layer/list", &controllers.LayuiadminGOController{},"*:ComponentLayerList")
	beego.Router("/component/layer/special-demo", &controllers.LayuiadminGOController{},"*:ComponentLayerSpecial_demo")
	beego.Router("/component/layer/theme", &controllers.LayuiadminGOController{},"*:ComponentLayerTheme")

	beego.Router("/component/laypage", &controllers.LayuiadminGOController{},"*:ComponentLaypageIndex")
	beego.Router("/component/laytpl", &controllers.LayuiadminGOController{},"*:ComponentLaytplIndex")
	beego.Router("/component/nav/index", &controllers.LayuiadminGOController{},"*:ComponentNavIndex")
	beego.Router("/component/panel", &controllers.LayuiadminGOController{},"*:ComponentPanelIndex")
	beego.Router("/component/progress", &controllers.LayuiadminGOController{},"*:ComponentProgressIndex")
	beego.Router("/component/rate", &controllers.LayuiadminGOController{},"*:ComponentRateIndex")
	beego.Router("/component/slider", &controllers.LayuiadminGOController{},"*:ComponentSliderIndex")


	///component laypage
	beego.Router("/component/laypage", &controllers.LayuiadminGOController{},"*:ComponentLaypageIndex")
	beego.Router("/component/laypage/index", &controllers.LayuiadminGOController{},"*:ComponentLaypageIndex")
	beego.Router("/component/laypage/demo1", &controllers.LayuiadminGOController{},"*:ComponentLaypageDemo1")
	beego.Router("/component/laypage/demo2", &controllers.LayuiadminGOController{},"*:ComponentLaypageDemo2")

	beego.Router("/component/laytpl/index", &controllers.LayuiadminGOController{},"*:ComponentLaytplIndex")
	beego.Router("/component/nav/index/index", &controllers.LayuiadminGOController{},"*:ComponentNavIndex")
	beego.Router("/component/panel/index", &controllers.LayuiadminGOController{},"*:ComponentPanelIndex")
	beego.Router("/component/progress/index", &controllers.LayuiadminGOController{},"*:ComponentProgressIndex")
	beego.Router("/component/rate/index", &controllers.LayuiadminGOController{},"*:ComponentRateIndex")
	beego.Router("/component/slider/index", &controllers.LayuiadminGOController{},"*:ComponentSliderIndex")

	//component table
	beego.Router("/component/table/auto", &controllers.LayuiadminGOController{},"*:ComponentTableAuto")
	beego.Router("/component/table/cellEdit", &controllers.LayuiadminGOController{},"*:ComponentTableCellEdit")
	beego.Router("/component/table/cellEvent", &controllers.LayuiadminGOController{},"*:ComponentTableCellEvent")
	beego.Router("/component/table/checkbox", &controllers.LayuiadminGOController{},"*:ComponentTableCheckbox")
	beego.Router("/component/table/data", &controllers.LayuiadminGOController{},"*:ComponentTableData")
	beego.Router("/component/table/fixed", &controllers.LayuiadminGOController{},"*:ComponentTableFixed")
	beego.Router("/component/table/form", &controllers.LayuiadminGOController{},"*:ComponentTableForm")
	beego.Router("/component/table/height", &controllers.LayuiadminGOController{},"*:ComponentTableHeight")
	beego.Router("/component/table/index", &controllers.LayuiadminGOController{},"*:ComponentTableIndex")
	beego.Router("/component/table/initSort", &controllers.LayuiadminGOController{},"*:ComponentTableInitSort")
	beego.Router("/component/table/onrow", &controllers.LayuiadminGOController{},"*:ComponentTableOnrow")
	beego.Router("/component/table/operate", &controllers.LayuiadminGOController{},"*:ComponentTableOperate")
	beego.Router("/component/table/page", &controllers.LayuiadminGOController{},"*:ComponentTablePage")
	beego.Router("/component/table/parseData", &controllers.LayuiadminGOController{},"*:ComponentTableParseData")
	beego.Router("/component/table/radio", &controllers.LayuiadminGOController{},"*:ComponentTableRadio")
	beego.Router("/component/table/reload", &controllers.LayuiadminGOController{},"*:ComponentTableReload")
	beego.Router("/component/table/resetPage", &controllers.LayuiadminGOController{},"*:ComponentTableResetPage")
	beego.Router("/component/table/simple", &controllers.LayuiadminGOController{},"*:ComponentTableSimple")
	beego.Router("/component/table/static", &controllers.LayuiadminGOController{},"*:ComponentTableStatic")
	beego.Router("/component/table/style", &controllers.LayuiadminGOController{},"*:ComponentTableStyle")
	beego.Router("/component/table/thead", &controllers.LayuiadminGOController{},"*:ComponentTableThead")
	beego.Router("/component/table/toolbar", &controllers.LayuiadminGOController{},"*:ComponentTableToolbar")
	beego.Router("/component/table/tostatic", &controllers.LayuiadminGOController{},"*:ComponentTableTostatic")
	beego.Router("/component/table/totalRow", &controllers.LayuiadminGOController{},"*:ComponentTableTotalRow")


	beego.Router("/component/tabs", &controllers.LayuiadminGOController{},"*:ComponentTabsIndex")
	beego.Router("/component/timeline", &controllers.LayuiadminGOController{},"*:ComponentTimelineIndex")

	beego.Router("/component/tabs/index", &controllers.LayuiadminGOController{},"*:ComponentTabsIndex")
	beego.Router("/component/timeline/index", &controllers.LayuiadminGOController{},"*:ComponentTimelineIndex")

	//component upload
	beego.Router("/component/upload", &controllers.LayuiadminGOController{},"*:ComponentUploadIndex")
	beego.Router("/component/upload/demo1", &controllers.LayuiadminGOController{},"*:ComponentUploadDemo1")
	beego.Router("/component/upload/demo2", &controllers.LayuiadminGOController{},"*:ComponentUploadDemo2")

	beego.Router("/component/util", &controllers.LayuiadminGOController{},"*:ComponentUtilIndex")
	beego.Router("/component/util/index", &controllers.LayuiadminGOController{},"*:ComponentUtilIndex")

}
