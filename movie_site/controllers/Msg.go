package controllers

import (
	"github.com/astaxie/beego"
	"SB/movie_site/services"
	"fmt"
)

type MsgController struct {
	beego.Controller
}

func (this *MsgController)OfficialMsg(){
	var resp services.MovieResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,err := this.GetInt("page")
	if err != nil {
		resp.Msg = "解析参数失败"
		return
	}
	resp = services.FindOfficalMsg(page)
	fmt.Println(resp.Next)
	return
}
