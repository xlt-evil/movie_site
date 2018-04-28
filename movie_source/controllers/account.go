package controllers

import (
	"github.com/astaxie/beego"
	"movie_source/service"
)


type AccountController struct {
	beego.Controller
}

//登录页面
func (this *AccountController)ToLogin(){
	this.TplName = "login.html"
}
//登录
func (this *AccountController)Login(){
	var resp service.MsgResp
	resp.Status = false
	defer this.ServeJSON()
	uid := this.GetString("user")
	password := this.GetString("password")
	resp = service.Login(uid,password)
	if resp.Status {
		this.Ctx.SetCookie("user",uid)
		this.Ctx.SetCookie("password",password )
	}
	this.Data["json"] = resp
}
//登出
func (this *AccountController)LoginOut(){
	uid := this.Ctx.GetCookie("uid")
	//清除redis
	service.LoginOut(uid)
	//重置cookie
	this.Ctx.SetCookie("user", "0", -1)
	this.Ctx.SetCookie("passwprd", "0", -1)
	//重定向到首页
	this.Ctx.Redirect(302,"/tologin")
}

