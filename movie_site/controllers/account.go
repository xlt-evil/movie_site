package controllers

import (
	"github.com/astaxie/beego"
	"SB/movie_site/models"
	"SB/movie_site/services"
	"fmt"
)

type AccountController struct{
	beego.Controller
}

//注册
func (this *AccountController)Register(){
	var u models.User
	defer this.ServeJSON()
	u.Uid = this.GetString("Uid")
	u.Question = this.GetString("Question")
	u.Answer = this.GetString("Answer")
	u.Password = this.GetString("Password")
	resp := services.AddUser(u)
	this.Data["json"] = resp
}

//登录
func (this *AccountController)Login(){
	var u models.User
	defer this.ServeJSON()
	u.Uid = this.GetString("Uid")
	u.Password = this.GetString("Password")
	resp := services.Login(u)
	if resp.Status {
		this.Ctx.SetCookie("uid",u.Uid)
		this.Ctx.SetCookie("pwd",u.Password )
	}
	this.Data["json"] = resp
}

//进入登录注册页面
func (this *AccountController) ToRegister(){
	this.TplName = "register.html"
}

//忘记密码
func (this *AccountController)ForgetPassword(){
	var u models.User
	defer this.ServeJSON()
	u.Uid = this.GetString("account")
	u.Answer = this.GetString("answer")
	u.Question = this.GetString("question")
	resp := services.FindPassword(u)
	this.Data["json"] = resp
}

//登出
func (this *AccountController)LoginOut(){
	uid := this.Ctx.GetCookie("uid")
	//清除redis
	services.LoginOut(uid)
	//重置cookie
	fmt.Print("gogogogogog")
	this.Ctx.SetCookie("uid", "0", -1)
	this.Ctx.SetCookie("pwd", "0", -1)
	//重定向到首页
	this.Ctx.Redirect(302,"/")
}



