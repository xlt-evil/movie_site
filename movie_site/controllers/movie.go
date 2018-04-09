package controllers

import (
	"github.com/astaxie/beego"
	"SB/movie_site/services"
	"fmt"
)

type MovieController struct {
	beego.Controller
}
//热门电影异步加载
func (this *MovieController)HotMovie(){
	condition := ""
	paras := []interface{}{}
	var resp services.MsgResponse
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	s:= this.GetStrings("country")
	if len(s) >0 {

	}
	for i,v := range  s {
		if i == 0{
			condition =" WHERE area = ? "
			paras = append(paras,v)
			continue
		}
		condition +=" or area = ? "
		paras = append(paras,v)
	}
	resp = services.FindMovieListByIndexAjax(condition,paras)
	return
}
//电影库异步加载
func (this *MovieController)MovieLibrary(){
	condition := ""
	order_by :=""
	var resp services.MovieResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	area,err := this.GetInt("area")
	if err != nil {
		resp.Msg = "参数获取失败area"
		return
	}
	if area != 0{
		condition += " AND area = ?"
	}
	types,err := this.GetInt("type")
	if err != nil {
		resp.Msg = "参数获取失败type"
		return
	}
	if types != 0{
		condition += " AND type = ?"
	}
	year := this.GetString("year")
	if year == "" {
		resp.Msg = "参数获取失败year"
		return
	}
	if year != "全部年代" && year != "2007以前"{
		condition += " AND year = ?"
	}
	if year == "2007以前"{
		condition += " AND year <= 2007"
	}
	sort,err := this.GetInt("sort")
	if err != nil {
		resp.Msg = "参数获取失败"
		return
	}
	page,err:= this.GetInt("page")
	if err != nil {
		resp.Msg = "参数获取失败"
		return
	}
	switch sort {
	case 1:
		order_by = " ORDER BY ma.popularity DESC,ma.id"
	case 2:
		order_by = " ORDER BY ma.collect_num DESC,ma.id"
	case 3:
		order_by = " ORDER BY ma.score DESC,ma.id"
	}
	resp = services.MovieLibraryAjax(area,types,page,year,condition,order_by)
	fmt.Println(area,types,page,year)
	return
}
//电影短评异步加载
func (this *MovieController)MovieTalkAjax(){
	var resp services.MovieResp
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page ,err := this.GetInt("page")
	if err!= nil {
		resp.Msg = "参数解析失败"
		return
	}
	movieId,err:=this.GetInt("movieId")
	if err != nil {
		resp.Msg = "参数解析失败"
		return
	}
	resp = services.MovieTalkAjax(movieId,page)
	return
}
//异步加载影评
func (this *MainController)ReviewAjaxList(){
	var resp services.MovieResp
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,err := this.GetInt("page")
	if err != nil{
		resp.Msg = "参数解析失败"
		return
	}
	resp = services.FindFilmReview(page)
	return
}
//异步填充页面
func(this *MainController)ShowTalkAjax(){
	var resp services.MsgResponse
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,_ :=this.GetInt("id")
	resp = services.TalkDetail(id)
	return
}
//影评评论异步加载
func (this *MainController)ReviewTalk(){
	var resp services.MovieResp
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,err := this.GetInt("page")
	if err != nil {
		resp.Msg = "解析参数失败"
		return
	}
	frid,err := this.GetInt("frid")
	if err != nil {
		resp.Msg = "解析参数失败"
		return
	}
	resp = services.LoadReleaseList(page,frid)
	return
}