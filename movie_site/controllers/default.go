package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"SB/movie_site/models"
	"SB/movie_site/services"
	"SB/movie_site/util"
	"fmt"
	"path"
	"os"
	"io"
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
		respsss := services.InsertHistory(this.User.Uid,id)//添加观看记录
		fmt.Println("历史记录",respsss)
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
	talk := services.LoadReleaseList(1,id)
	this.Data["list"] = talk
	this.TplName = "showTalk.html"
}
//资讯页面
func(this *MainController)News(){
	//resp := services.FindOfficalMsg(1)
	//this.Data["official"] = resp
	this.TplName = "news.html"
}











func(this *MainController)Demo(){
	this.TplName = "demo.html"
}
//测试demo
func(this *MainController)Upload(){
	fmt.Println("123asdas")
	var resp services.MsgResponse
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	s,f,e :=this.GetFile("file")
	if e != nil {
		resp.Msg = "文件获取失败"+e.Error()
		return
	}

	fmt.Println(f.Filename)
	str := path.Join("http://localhost:9090/static/css",f.Filename)
	fmt.Println(str)
	err := this.SaveToFile("image", "static/css/1.jpg")
	if err != nil {
		resp.Msg = err.Error()
	}
	f1,e := os.Create("static/css/1.jpg")
	io.Copy(f1,s)
	s.Close()
}




