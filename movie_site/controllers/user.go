package controllers

import (
	"SB/movie_site/services"
	"SB/movie_site/models"
)

type UserController struct {
	BaseController
}
//个人中心
func (this *UserController)ToPerson(){
	if this.User == nil {
		this.Ctx.Redirect(302,"/")
	}
	this.TplName = "person.html"
}

//个人电影收藏
func (this *UserController)CollectFilm(){
	var resp services.MsgResponse
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登录"
		return
	}
	movieId,err:= this.GetInt("movieId")
	if err != nil {
		resp.Msg = "参数解析失败"
		return
	}

	collect := this.GetString("status")
	resp = services.CollecFilm(this.User.Uid,movieId,collect)
	return
}

//个人评分
func (this *UserController)ScoreFilm(){
	var resp services.MsgResponse
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登录"
		return
	}
	movieId,err:= this.GetInt("movieId")
	if err != nil {
		resp.Msg = "参数解析失败"
		return
	}

	score ,err:= this.GetInt("score")
	if err != nil {
		resp.Msg = "参数解析失败"
		return
	}
	resp = services.ScoreFilm(this.User.Uid,movieId,score)
	return
}

func(this *UserController)MovieTalk(){
	var resp services.MsgResponse
	var mt models.MovieTalk
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	movieId,err:= this.GetInt("movieId")
	if err != nil {
		resp.Msg = "参数解析失败"
		return
	}
	mt.MovieId = movieId
	content := this.GetString("content")
	mt.Content = content
	if this.User == nil {
		resp.Msg = "请登入"
		return
	}
	mt.UserId = this.User.Uid
	resp = services.MovieTalk(&mt)
	resp.Status = true
	return
}

//影评页面
func(this *UserController)FilmCritic(){
	if this.User == nil {
		this.Redirect("/",302)
	}
	id,err:= this.GetInt("movieId")
	if err != nil {
		this.Abort("302")
	}
	name,id:= services.FindMovieName(id)
	this.Data["name"] = name
	this.Data["id"] = id
	this.TplName = "write.html"
}

//发布影评
func (this *UserController)ReleaseFilmCritic(){
	var resp services.MsgResponse
	var fr models.FilmReview
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "用户不存在！"
		return
	}
	fr.UserId = this.User.Uid
	content := this.GetString("content")
	if content == ""{
		resp.Msg = "提交的内容为空，发布失败！"
		return
	}
	fr.Content = content
	id,err := this.GetInt("movieId")
	if err != nil{
		resp.Msg = "参数解析失败"
		return
	}
	fr.MovieId = id
	title := this.GetString("title")
	if title == "" {
		resp.Msg ="标题为空提交失败"
		return
	}
	fr.Title = title
	resp = services.ReleaseFilmCritic(&fr)
	return
}
//影评喜好
func (this *UserController)FilmReviewFeel(){
	var resp services.MsgResponse
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登录"
		return
	}
	feel,err := this.GetInt("feel")
	if err != nil {
		resp.Msg = "解析参数失败"
		return
	}
	frid,err := this.GetInt("frid")
	if err != nil{
		resp.Msg = "解析参数失败"
		return
	}
	resp = services.ReviewFeel(frid,this.User.Uid,feel)
	return
}




