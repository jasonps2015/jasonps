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

	//iframe
	beego.Router("/Iframe", &layuiadmingo.LayuiadminGOController{},"*:IframeLayerIframe")
	beego.Router("/Iframe/Layer/Iframe", &layuiadmingo.LayuiadminGOController{},"*:IframeLayerIframe")


	//system
	beego.Router("/system", &layuiadmingo.LayuiadminGOController{},"*:SystemAbout")
	beego.Router("/system/about", &layuiadmingo.LayuiadminGOController{},"*:SystemAbout")
	beego.Router("/system/get", &layuiadmingo.LayuiadminGOController{},"*:SystemGet")
	beego.Router("/system/more", &layuiadmingo.LayuiadminGOController{},"*:SystemMore")
	beego.Router("/system/theme", &layuiadmingo.LayuiadminGOController{},"*:SystemTheme")

	//component
	beego.Router("/component/", &layuiadmingo.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/anim", &layuiadmingo.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/auxiliar", &layuiadmingo.LayuiadminGOController{},"*:ComponentAuxiliarIndex")
	beego.Router("/component/badge", &layuiadmingo.LayuiadminGOController{},"*:ComponentBadgeIndex")
	beego.Router("/component/button", &layuiadmingo.LayuiadminGOController{},"*:ComponentButtonIndex")
	beego.Router("/component/carousel", &layuiadmingo.LayuiadminGOController{},"*:ComponentCarouselIndex")
	beego.Router("/component/code", &layuiadmingo.LayuiadminGOController{},"*:ComponentCodeIndex")
	beego.Router("/component/colorpicker", &layuiadmingo.LayuiadminGOController{},"*:ComponentColorpickerIndex")
	beego.Router("/component/flow", &layuiadmingo.LayuiadminGOController{},"*:ComponentFlowIndex")

	beego.Router("/component/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/anim/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentAnimIndex")
	beego.Router("/component/auxiliar/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentAuxiliarIndex")
	beego.Router("/component/badge/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentBadgeIndex")
	beego.Router("/component/button/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentButtonIndex")
	beego.Router("/component/carousel/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentCarouselIndex")
	beego.Router("/component/code/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentCodeIndex")
	beego.Router("/component/colorpicker/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentColorpickerIndex")
	beego.Router("/component/flow/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentFlowIndex")

	//component form
	beego.Router("/component/form", &layuiadmingo.LayuiadminGOController{},"*:ComponentFormElement")
	beego.Router("/component/form/element", &layuiadmingo.LayuiadminGOController{},"*:ComponentFormElement")
	beego.Router("/component/form/group", &layuiadmingo.LayuiadminGOController{},"*:ComponentFormGroup")
	//component grid
	beego.Router("/component/grid", &layuiadmingo.LayuiadminGOController{},"*:ComponentGridAll")
	beego.Router("/component/grid/all", &layuiadmingo.LayuiadminGOController{},"*:ComponentGridAll")
	beego.Router("/component/grid/list", &layuiadmingo.LayuiadminGOController{},"*:ComponentGridList")
	beego.Router("/component/grid/mobile", &layuiadmingo.LayuiadminGOController{},"*:ComponentGridMobile")
	beego.Router("/component/grid/mobile-pc", &layuiadmingo.LayuiadminGOController{},"*:ComponentGridMobile_pc")
	beego.Router("/component/grid/speed-dial", &layuiadmingo.LayuiadminGOController{},"*:ComponentGridSpeed_dial")
	beego.Router("/component/grid/stack", &layuiadmingo.LayuiadminGOController{},"*:ComponentGridStack")



	//component laydate
	beego.Router("/component/laydate", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaydateIndex")
	beego.Router("/component/laydate/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaydateIndex")
	beego.Router("/component/laydate/demo1", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaydateDemo1")
	beego.Router("/component/laydate/demo2", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaydateDemo2")
	beego.Router("/component/laydate/special-demo", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaydateSpecial_demo")
	beego.Router("/component/laydate/theme", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaydateTheme")

	//beego.Router("/component/layer", &layuiadmingo.LayuiadminGOController{},"*:ComponentLayerList")
	beego.Router("/component/layer/list", &layuiadmingo.LayuiadminGOController{},"*:ComponentLayerList")
	beego.Router("/component/layer/special-demo", &layuiadmingo.LayuiadminGOController{},"*:ComponentLayerSpecial_demo")
	beego.Router("/component/layer/theme", &layuiadmingo.LayuiadminGOController{},"*:ComponentLayerTheme")

	beego.Router("/component/laypage", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaypageIndex")
	beego.Router("/component/laytpl", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaytplIndex")
	beego.Router("/component/nav/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentNavIndex")
	beego.Router("/component/panel", &layuiadmingo.LayuiadminGOController{},"*:ComponentPanelIndex")
	beego.Router("/component/progress", &layuiadmingo.LayuiadminGOController{},"*:ComponentProgressIndex")
	beego.Router("/component/rate", &layuiadmingo.LayuiadminGOController{},"*:ComponentRateIndex")
	beego.Router("/component/slider", &layuiadmingo.LayuiadminGOController{},"*:ComponentSliderIndex")


	///component laypage
	beego.Router("/component/laypage", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaypageIndex")
	beego.Router("/component/laypage/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaypageIndex")
	beego.Router("/component/laypage/demo1", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaypageDemo1")
	beego.Router("/component/laypage/demo2", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaypageDemo2")

	beego.Router("/component/laytpl/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentLaytplIndex")
	beego.Router("/component/nav/index/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentNavIndex")
	beego.Router("/component/panel/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentPanelIndex")
	beego.Router("/component/progress/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentProgressIndex")
	beego.Router("/component/rate/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentRateIndex")
	beego.Router("/component/slider/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentSliderIndex")

	//component table
	beego.Router("/component/table/auto", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableAuto")
	beego.Router("/component/table/cellEdit", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableCellEdit")
	beego.Router("/component/table/cellEvent", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableCellEvent")
	beego.Router("/component/table/checkbox", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableCheckbox")
	beego.Router("/component/table/data", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableData")
	beego.Router("/component/table/fixed", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableFixed")
	beego.Router("/component/table/form", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableForm")
	beego.Router("/component/table/height", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableHeight")
	beego.Router("/component/table/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableIndex")
	beego.Router("/component/table/initSort", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableInitSort")
	beego.Router("/component/table/onrow", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableOnrow")
	beego.Router("/component/table/operate", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableOperate")
	beego.Router("/component/table/page", &layuiadmingo.LayuiadminGOController{},"*:ComponentTablePage")
	beego.Router("/component/table/parseData", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableParseData")
	beego.Router("/component/table/radio", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableRadio")
	beego.Router("/component/table/reload", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableReload")
	beego.Router("/component/table/resetPage", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableResetPage")
	beego.Router("/component/table/simple", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableSimple")
	beego.Router("/component/table/static", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableStatic")
	beego.Router("/component/table/style", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableStyle")
	beego.Router("/component/table/thead", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableThead")
	beego.Router("/component/table/toolbar", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableToolbar")
	beego.Router("/component/table/tostatic", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableTostatic")
	beego.Router("/component/table/totalRow", &layuiadmingo.LayuiadminGOController{},"*:ComponentTableTotalRow")


	beego.Router("/component/tabs", &layuiadmingo.LayuiadminGOController{},"*:ComponentTabsIndex")
	beego.Router("/component/timeline", &layuiadmingo.LayuiadminGOController{},"*:ComponentTimelineIndex")

	beego.Router("/component/tabs/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentTabsIndex")
	beego.Router("/component/timeline/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentTimelineIndex")

	//component upload
	beego.Router("/component/upload", &layuiadmingo.LayuiadminGOController{},"*:ComponentUploadIndex")
	beego.Router("/component/upload/demo1", &layuiadmingo.LayuiadminGOController{},"*:ComponentUploadDemo1")
	beego.Router("/component/upload/demo2", &layuiadmingo.LayuiadminGOController{},"*:ComponentUploadDemo2")

	beego.Router("/component/util", &layuiadmingo.LayuiadminGOController{},"*:ComponentUtilIndex")
	beego.Router("/component/util/index", &layuiadmingo.LayuiadminGOController{},"*:ComponentUtilIndex")

}
