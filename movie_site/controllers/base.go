package controllers

import (
	"github.com/astaxie/beego"
	"movie_site/models"
	"movie_site/util"
	"encoding/json"
)

//以登录用户为基础的controller
type BaseController struct{
	beego.Controller
	User *models.User
}


//判断用户是否在线
func (this *BaseController)Prepare(){
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
	}else{
		this.Ctx.Redirect(302,"/")
	}
}


