package services

import (
	"movie_site/models"
	"strconv"
	"movie_site/util"
	"encoding/json"
	"time"
)

//信息返回体
type MsgResponse struct {
	Status bool   `description:"通过状态"`
	Msg    string `description:"信息"`
	Object interface{} `description:"传值的"`
}

//添加用户
func AddUser(u models.User)(resp MsgResponse){
	name := models.CheckUid(u.Uid)
	if name != "" {
		resp.Status = false
		resp.Msg = "用户已存在"
		return
	}else{
		err := models.AddUser(u)
		if err != nil {
			resp.Status = false
			resp.Msg = "发生未知的异常请重试"
			return
		}
		resp.Status = true
		resp.Msg = "注册成功"
	}
	return
}


//登录
func Login(u models.User)(resp MsgResponse){
	user := models.CheckPwdAndUid(u)
	resp.Status = false
	if user == nil {
		resp.Msg = "账户或密码错误请重试！"
		return
	}
	user.ImgHead = util.Person_img + user.ImgHead
	b ,_ := json.Marshal(&user)
	//登录成功的用户添加到缓存区域
	if util.Re == nil{
		util.RedisPut(util.Rc,util.BaseKeyName+user.Uid,b,24*time.Hour)
	}
	resp.Status = true
	resp.Msg = "登录成功，欢迎"+strconv.Itoa(user.Id)+"的到来"
	return
}

//登出
func LoginOut(key string)(resp MsgResponse){
	if util.Re == nil{
		util.DelKey(util.Rc,util.BaseKeyName+key)
		resp.Status = true
		resp.Msg = "欢迎下次光临！"
	}else{
		resp.Status = false
		resp.Msg = "服务器异常，请稍后重试"
	}
	return
}

func FindPassword(u models.User)(resp MsgResponse){
	password := models.ForgetPwd(u)
	if password == "" {
		resp.Status = false
		resp.Msg = "你的回答或账号错误！"
		return
	}
	resp.Status = true
	resp.Msg = "你的密码是："
	resp.Object = password
	return
}
