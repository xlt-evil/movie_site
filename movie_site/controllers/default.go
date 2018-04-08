package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"SB/movie_site/models"
	"SB/movie_site/services"
	"SB/movie_site/util"
)

//一般的页面跳转 ，以及首次页面填充

type MainController struct {
	beego.Controller
	User *models.User
}


func (this *MainController) Prepare() {
	uid := this.Ctx.GetCookie("uid")
	pwd := this.Ctx.GetCookie("pwd")
	if uid != "" && pwd != "" {
		if util.Re == nil{
			if util.IsExist(util.Rc,util.BaseKeyName+uid) {
				if data ,err :=util.RedisRead(util.Rc,util.BaseKeyName+uid);err==nil{
					err = json.Unmarshal(data,&this.User)
					this.Data["user"] = this.User
				}
			}
		}
	}
}

//首页
func (this *MainController)Get(){
	key := make([]string,2)
	key[0] = "日本"
	key[1] = "韩国"
	resp :=services.FindMovieListByIndex(key)
	this.Data["movies"] = resp
	this.TplName = "index.html"
}

//电影库
func (this *MainController)ToMovieLibrary(){
	resp := services.FindAreaMenu()
	resp1 := services.FindTypeMenu()
	resp2 := services.MovieLibrary()
	this.Data["resp"] = resp
	this.Data["resp1"] = resp1
	this.Data["resp2"] = resp2
	this.TplName = "movielibrary.html"
}

//电影介绍页面
func (this *MainController)MovieIntroduce(){
	id,err := this.GetInt("movieId")
	resp := services.FindMovie(id)
	resp1 := services.FindMovieAttribute(id)
	movieTalk := services.MovieTalkAjax(id,1)
	if this.User == nil {
		this.Data["person"] = models.PersonFilm{Collect:"false",Score:0}
	}else{
		user := services.CollectStatus(this.User.Uid,id)
		this.Data["person"] = user.Object
	}
	if err != nil {
		resp.Msg = "获取参数失败"
	}
	this.Data["movie"] = resp
	this.Data["attribute"] = resp1
	this.Data["talk"] = movieTalk
	this.TplName = "movieintrdouce.html"
}

//电影放映厅
func (this *MainController)MoviePlayer(){
	id,err := this.GetInt("movieId")
	resp := services.FindMovie(id)
	resp1 := services.FindMovieAttribute(id)
	movieTalk := services.MovieTalkAjax(id,1)
	if this.User == nil {
		this.Data["person"] = models.PersonFilm{Collect:"false",Score:0}
	}else{
		user := services.CollectStatus(this.User.Uid,id)
		this.Data["person"] = user.Object
	}
	if err != nil {
		resp.Msg = "获取参数失败"
	}
	this.Data["movie"] = resp
	this.Data["attribute"] = resp1
	this.Data["talk"] = movieTalk
	this.TplName = "movie.html"
}

//电影影评
func(this *MainController)FilmReview(){
	resp := services.FindFilmReview(1)
	this.Data["review"] = resp
	this.TplName = "movieTalklist.html"
}

//影评的具体内容
func (this *MainController)ShowTalk(){
	id,err:= this.GetInt("id")
	if err != nil {
		this.Abort("404")
	}
	resp := services.TalkDetail(id)
	this.Data["talks"] = resp.Object
	this.TplName = "showTalk.html"
}


