package controllers

import (
	"movie_source/service"
	"strconv"
)

//用户管理
type UserController struct {
	BaseController
}

//首页
func (this *UserController)Index(){
	this.IsNeedTemplate()
	resp := service.CheckUserDate()
	resp1 := service.AdviseList(1)
	resp2 := service.ResourceList(1)
	this.Data["Title"] = "用户管理"
	this.Data["title"] = "user-manager"
	for i,_ := range  resp {
		this.Data["count"+strconv.Itoa(i)] = resp[i]
	}
	this.Data["advise"] = resp1
 	this.Data["resource"] = resp2
	this.TplName = "user/index.html"
}

//用户消息列表查看
func (this *UserController)UserList(){
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
	resp = service.AdviseList(page)
}

//建设性消息
func (this *UserController)ResourceList(){
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
	resp = service.ResourceList(page)
}

//消息具体信息
func (this *UserController)Detail(){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id ,err := this.GetInt("id")
	if err != nil{
		resp.Msg = "解析参数失败"
		return
	}
	resp = service.MsgDetail(id)
}