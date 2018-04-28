package controllers

import (
	"movie_source/service"
	"movie_source/util"
	"fmt"
	"time"
)

//资讯管理
type InformationController struct {
	BaseController
}

//新闻页面
func (this *InformationController)News(){
	this.IsNeedTemplate()
	resp := service.NewsList(1)
	this.Data["Title"] = "新闻管理"
	this.Data["title"] = "news-manage"
	this.Data["news"] = resp
	this.TplName = "information/news.html"
}

//发布新闻
func (this *InformationController)CreateNews(){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	title := this.GetString("title")
	point := this.GetString("point")
	content := this.GetString("content")
	id := int(time.Now().Unix())
	s,f,err := this.GetFile("img")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	name := util.SaveImg(s,f,id,"news_img/")
	s.Close()
	resp = service.ReleaseNews(title,point,content,name)
	return
}

//用户消息列表查看
func (this *InformationController)NewsList(){
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
	resp = service.NewsList(page)
}
//删除新闻
func (this *InformationController)DeteleNews(){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,err := this.GetInt("id")
	if err != nil {
		resp.Msg = "解析参数失败"
		return
	}
	resp = service.DelNews(id)
	return

}

//官方公告
func (this *InformationController)Offical(){
	this.IsNeedTemplate()
	resp := service.OfficalList(1)
	this.Data["msg"] = resp
	fmt.Println(resp.Object)
	this.Data["Title"] = "官方消息"
	this.Data["title"] = "offical-msg"
	this.TplName = "information/offical.html"
}

//官方消息列表
func (this *InformationController)OfficialList(){
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
	resp = service.OfficalList(page)
	return
}
//发布消息
func (this *InformationController)ReleaseOfficialMsg(){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	title := this.GetString("title")
	content := this.GetString("content")
	resp = service.ReleaseOfficialMsg(title,content)
	return
}
//删除消息
func (this *InformationController)DelOfficialMsg(){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,err := this.GetInt("id")
	if err != nil{
		resp.Msg = "解析参数失败"
		return
	}
	resp = service.DelOfficialMsg(id)
}