package controllers

import (
	"movie_source/service"
	"fmt"
	"movie_source/util"
	"time"
)

//电影管理
type MovieController struct {
	BaseController
}

//电影管理
func (this *MovieController) MovieManage(){
	this.IsNeedTemplate()
	resp := service.MovieList(1)
	t,c := service.MovieCount()
	a,ts,l:= service.Menu()
	this.Data["Title"] = "电影管理"
	this.Data["title"] = "movie-manage"
	this.Data["movie"] = resp
	this.Data["t"] = t
	this.Data["c"] = c
	this.Data["area"] = a
	this.Data["type"] = ts
	this.Data["language"] = l
	this.TplName = "movie/moviemanage.html"
}

//电影发布
func (this *MovieController) MovieRelease(){
	this.IsNeedTemplate()
	a,ts,l:= service.Menu()
	this.Data["Title"] = "电影发布"
	this.Data["title"] = "movie-release"
	this.Data["area"] = a
	this.Data["type"] = ts
	this.Data["language"] = l
	this.TplName = "movie/movierelease.html"
}

//电影分页
func (this *MovieController)MovieList(){
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
	resp = service.MovieList(page)
	return
}

//查看电影详情
func (this *MovieController)FindMovie(){
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
	fmt.Println(id)
	resp = service.GetMovieDetail(id)
	return
}
//修改电影资料
func (this *MovieController) ModifyMovie (){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	var id = 0
	var img = ""
	status ,_ := this.GetInt("status",0)
	if status != 0 {
		s,f,e := this.GetFile("img")
		if e != nil {
			resp.Msg = e.Error()
			return
		}
		id = int(time.Now().Unix())
		img = util.SaveImg(s,f,int(time.Now().Unix()),"movie_img/")
	}
	if id == 0 && status != 0{
		resp.Msg = "保存图片错误"
		return
	}
	movieId,_ := this.GetInt("id")
	moviename := this.GetString("moviename")
	director := this.GetString("director")
	area := this.GetString("area")
	types := this.GetString("type")
	language := this.GetString("language")
	releasedate := this.GetString("releasedate")
	content := this.GetString("content")
	short := this.GetString("short")
	resp = service.ModifyMovie(movieId,moviename,director,area,types,language,releasedate,content,short,img,status)
	fmt.Println(resp)
	return
}
//电影上传
func (this *MovieController)MovieUpload(){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		if (resp.Status == false){
			this.Ctx.WriteString("上传失败")
		}
		this.Redirect("./ok",302)
	}()
	movieId ,err := this.GetInt("movieId")
	if err != nil {
		resp.Msg = "上传失败，id不存在！"
		return
	}
	s,f,e := this.GetFile("file")
	if e != nil {
		resp.Msg = e.Error()
		return
	}
	id := int(time.Now().Unix())
	path := util.SaveVideo(s,f,id)
	resp = service.UpdateMovieSource(path,movieId)
	return
}

func (this *MovieController)ReleaseMovie(){
	var resp service.MsgResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	var id = 0
	var img = ""
	var path = ""
	status ,_ := this.GetInt("status",0)
	if status != 0 {
		s,f,e := this.GetFile("img")
		if e != nil {
			resp.Msg = e.Error()
			return
		}
		id = int(time.Now().Unix())
		img = util.SaveImg(s,f,int(time.Now().Unix()),"movie_img/")
	}
	if id == 0 && status != 0{
		resp.Msg = "保存图片错误"
		return
	}
	videostatus,_ := this.GetInt("videostatus",0)
	if videostatus != 0 {
		s,f,e := this.GetFile("video")
		if e != nil {
			resp.Msg = e.Error()
			return
		}
		id = int(time.Now().Unix())
		path = util.SaveVideo(s,f,id)
	}
	if id == 0 && videostatus != 0{
		resp.Msg = "电影上传失败"
		return
	}
	moviename := this.GetString("moviename")
	director := this.GetString("director")
	area := this.GetString("area")
	types := this.GetString("type")
	language := this.GetString("language")
	releasedate := this.GetString("releasedate")
	content := this.GetString("content")
	short := this.GetString("short")
	resp = service.ReleaseMovie(moviename,director,area,types,language,releasedate,content,short,img,status,videostatus,path)
}
func (this *MovieController)Ok(){
	this.TplName = "ok.html"
}


