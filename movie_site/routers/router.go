package routers

import (
	"movie_site/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//基本页面跳转使用固定路由
    beego.Router("/", &controllers.MainController{})
	beego.Router("/tomovielirbary", &controllers.MainController{},"get:ToMovieLibrary")
	beego.Router("/movieintrdouce", &controllers.MainController{},"get:MovieIntroduce")
	beego.Router("/movieplayer", &controllers.MainController{},"get:MoviePlayer")


    //登录注册模块
    beego.AutoRouter(&controllers.AccountController{})

	//用户模块
	beego.AutoRouter(&controllers.UserController{})
}
