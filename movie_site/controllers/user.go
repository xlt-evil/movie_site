package controllers

import (
	"SB/movie_site/services"
	"SB/movie_site/models"
	"fmt"
)

type UserController struct {
	BaseController
}
//个人中心
func (this *UserController)ToPerson(){
	if this.User == nil {
		this.Ctx.Redirect(302,"/")
	}
	resp := services.FindLikeMovieByUid(this.User.Uid,1)//收藏电影
	resp1 := services.FineMyReviewByUid(this.User.Uid,1)//影评
	resp2 := services.GetUserHistoryRecord(this.User.Uid,1)//观看记录
	this.Data["movie"] = resp
	this.Data["review"] = resp1
	this.Data["history"] = resp2
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
//电影短评发布
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
//发布影评评论
func (this *UserController)ReviewsRelease(){
	var resp services.MsgResponse
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登录"
		return
	}
	frid,err := this.GetInt("frid")
	if err != nil {
		resp.Msg = "参数解析出错"
		return
	}
	movieId,err := this.GetInt("movieId")
	if err != nil {
		resp.Msg = "参数解析出错"
		return
	}
	content := this.GetString("content")
	resp = services.ReviewsRelease(frid,movieId,this.User.Uid,content)
}
//查看个人电影收藏
func (this *UserController)CheckLikeMovie(){
	var resp services.MovieResp
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登录"
		return
	}
	page, err := this.GetInt("page")
	if err != nil {
		resp.Msg = "参数解析出错"
		return
	}
	resp = services.FindLikeMovieByUid(this.User.Uid,page)
	return
}
//查看个人影评喜好
func (this *UserController)CheckMyReview(){
	var resp services.MovieResp
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请登录"
		return
	}
	page, err := this.GetInt("page")
	if err != nil {
		resp.Msg = "参数解析出错"
		return
	}
	resp = services.FineMyReviewByUid(this.User.Uid,page)
	return
}
//个人基本信息修改
func (this *UserController)PersonMsgUpdate(){
	var resp services.MsgResponse
	resp.Status = true
	fmt.Println("dwfnadkjhfbaaaaaaaaaaaaaaaaaaaaaaaaaa")
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		this.Ctx.Redirect(302,"/")
		return
	}
	condition := ""
	paras := []string{}
	nick := this.GetString("nick")
	if nick != "" {
		condition += " name = ? ,"
		paras = append(paras,nick)
	}
	birth := this.GetString("birth")
	if birth != ""{
		condition += " birthday = ? ,"
		paras = append(paras,birth)
	}
	sign := this.GetString("sign")
	if sign != ""{
		condition += " person_sign = ? ,"
		paras = append(paras,sign)
	}
	if condition != ""{
		condition = condition[:len(condition)-2]
	}
	resp = services.UpdatePersonBaseMsg(condition,paras,this.User.Uid)
	return
}
//找到个人信息
func (this *UserController)UpdateMyNotice(){
	var resp services.MsgResponse
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		this.Ctx.Redirect(302,"/")
	}
	notice := this.GetString("notice")
	resp = services.UpdateMyNotice(this.User.Uid,notice)
	return
}
//更新密码
func (this *UserController)UpdatePassword(){
	var resp services.MsgResponse
	resp.Status = false
	defer func() {
		this.Data["json"]  = resp
		this.ServeJSON()
	}()
	old := this.GetString("old")
	if old == ""{
		resp.Msg = "密码不能为空"
		return
	}
	new := this.GetString("new")
	if new == ""{
		resp.Msg = "密码不能为空"
		return
	}
	resp = services.UpdateMyPassword(this.User.Uid,new,old)
	return
}
//更新头像
func (this *UserController)UpdateImg(){
	flag := this.GetString("flag")
	defer func() {
		this.Redirect("/user/toperson",302)
	}()
	if flag == "0" {
		return
	}
	uid := this.GetString("uid")
	imgname := this.GetString("img")
	services.UpdateImg(uid,imgname)
}
//查看观影历史
func (this *UserController)CheckHistoryRecord(){
	var resp services.MovieResp
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,err := this.GetInt("page")
	if err != nil {
		resp.Msg = "参数解析失败"
		return
	}
	fmt.Println("page",page)
	resp = services.GetUserHistoryRecord(this.User.Uid,page)
	return
}
//删除历史记录
func (this *UserController)DeleteHistoryRecord(){
	var resp services.MsgResponse
	resp.Status = false
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	movieId ,err:=this.GetInt("movieId")
	if err != nil {
		resp.Msg = "解析参数失败"
		return
	}
	resp = services.DeleteMyHistory(this.User.Uid,movieId)
	return
}


