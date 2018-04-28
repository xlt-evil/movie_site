package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"movie_source/util"
	"encoding/json"
	"movie_source/models"
)

type BaseController struct {
	User models.Admin
	beego.Controller
}

func (this *BaseController)Prepare(){
	uid := this.Ctx.GetCookie("user")
	pwd := this.Ctx.GetCookie("password")
	if uid != "" && pwd != "" {
		if util.Re == nil{
			if util.IsExist(util.Rc,util.BaseKeyName+uid) {
				if data ,err :=util.RedisRead(util.Rc,util.BaseKeyName+uid);err==nil{
					err = json.Unmarshal(data,&this.User)
					this.Data["user"] = this.User
				}
			}
		}
	} else{
		this.Ctx.Redirect(302,"/tologin")
	}
	fmt.Println(uid,pwd)

}



//模板
func (this *BaseController)IsNeedTemplate(){
	this.Layout = "layout.html"
}

