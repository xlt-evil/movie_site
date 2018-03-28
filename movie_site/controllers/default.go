package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"movie_site/util"
	"movie_site/models"
	"movie_site/services"
	"fmt"
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
	fmt.Print("resp",resp)
	this.Data["movies"] = resp
	this.TplName = "index.html"
}

//电影库
func (this *MainController)ToMovieLibrary(){
	resp := services.FindAreaMenu()
	resp1 := services.FindTypeMenu()
	this.Data["resp"] = resp
	this.Data["resp1"] = resp1
	this.TplName = "movielibrary.html"
}

//电影介绍页面
func (this *MainController)MovieIntroduce(){
	id,err := this.GetInt("movieId")
	resp := services.FindMovie(id)
	resp1 := services.FindMovieAttribute(id)
	if err != nil {
		resp.Msg = "获取参数失败"
	}
	this.Data["movie"] = resp
	this.Data["attribute"] = resp1
	this.TplName = "movieintrdouce.html"
}

//电影放映厅
func (this *MainController)MoviePlayer(){
	id,err := this.GetInt("movieId")
	resp := services.FindMovie(id)
	resp1 := services.FindMovieAttribute(id)
	if err != nil {
		resp.Msg = "获取参数失败"
	}
	this.Data["movie"] = resp
	this.Data["attribute"] = resp1
	this.TplName = "movie.html"
}

