package routers

import (
	"github.com/astaxie/beego"
	"movie_source/controllers"
)

func init() {
	//固定路由
	beego.Router("/updateImg",&controllers.MainController{},"post:UpdateImg")
	beego.Router("/tologin",&controllers.AccountController{},"get:ToLogin")
	beego.Router("/",&controllers.UserController{},"get:Index")

	//自动路由
	beego.AutoRouter(&controllers.BaseController{})
	beego.AutoRouter(&controllers.AccountController{})//登录管理
	beego.AutoRouter(&controllers.UserController{})//用户管理
	beego.AutoRouter(&controllers.MovieController{})//电影管理
	beego.AutoRouter(&controllers.InformationController{})//资讯管理
	beego.AutoRouter(&controllers.SystemController{})//系统管理

}
