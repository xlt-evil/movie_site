package routers

import (
	"SB/movie_site/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//基本页面跳转使用固定路由
    beego.Router("/", &controllers.MainController{})
	beego.Router("/tomovielirbary", &controllers.MainController{},"get:ToMovieLibrary")
	beego.Router("/movieintrdouce", &controllers.MainController{},"get:MovieIntroduce")
	beego.Router("/movieplayer", &controllers.MainController{},"get:MoviePlayer")
	beego.Router("/filmreview", &controllers.MainController{},"get:FilmReview")
	beego.Router("/reviewajaxlist", &controllers.MainController{},"get:ReviewAjaxList")
	beego.Router("/showtalk",&controllers.MainController{},"get:ShowTalk")
	beego.Router("/showtalkajax",&controllers.MainController{},"get:ShowTalkAjax")
	beego.Router("/reviewtalk",&controllers.MainController{},"get:ReviewTalk")
	beego.Router("/news",&controllers.MainController{},"get:News")
	beego.Router("/demo",&controllers.MainController{},"get:Demo")
	beego.Router("/upload",&controllers.MainController{},"post:Upload")


    //登录注册模块
    beego.AutoRouter(&controllers.AccountController{})

	//用户模块
	beego.AutoRouter(&controllers.UserController{})

	//电影模块
	beego.AutoRouter(&controllers.MovieController{})

	//消息模块
	beego.AutoRouter(&controllers.MsgController{})
}
