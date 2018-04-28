package controllers

import (
	"movie_source/service"
	"fmt"
)

//系统管理
type SystemController struct {
	BaseController
}

//系统用户管理
func (this *SystemController)SysUserManager(){
	this.IsNeedTemplate()
	resp := service.AdminList(1)
	fmt.Println(resp.Msg)
	this.Data["Title"] = "系统用户"
	this.Data["title"] = "system-user"
	this.Data["admin"] = resp
	this.TplName = "system/systemuser.html"
}
//添加用户页面
func (this *SystemController)SystemUser(){
	this.IsNeedTemplate()
	this.Data["Title"] = "添加用户"
	this.Data["title"] = "Add-systemuser"
	this.TplName = "system/adduser.html"
}

//用户列表
func (this *SystemController)SystemList(){
	var resp service.PagemsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,err := this.GetInt("page")
	if err != nil{
		resp.Msg = "解析参数失败"
		return
	}
	resp = service.AdminList(page)
}

//添加用户
func (this *SystemController)RegisterAdmin(){
	var resp service.MsgResp
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	role := this.User.Role
	uid := this.GetString("uid")
	password := this.GetString("password")
	resp = service.AddAdmin(role,uid,password)
	return
}

//冻结账户
func (this *SystemController)ForzenAdmin(){
	var resp service.MsgResp
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	role := this.User.Role
	id,err := this.GetInt("id")
	if err != nil{
		resp.Msg = "解析参数失败"
		return
	}
	status,err := this.GetInt("status")
	if err != nil{
		resp.Msg = "解析参数失败"
		return
	}
	resp = service.Frozen(role,id,status)
}

//删除用户
func (this *SystemController)DelAdmin(){
	var resp service.MsgResp
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	role := this.User.Role
	id,err := this.GetInt("id")
	if err != nil{
		resp.Msg = "解析参数失败"
		return
	}
	resp = service.DelAdmin(role,id)
	return
}