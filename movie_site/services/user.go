package services

import (
	"SB/movie_site/models"
	"strconv"
	"SB/movie_site/util"
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
//找到密码
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
//收藏电影
func CollecFilm(uid string,movieId int,collect string)(resp MsgResponse){
	resp.Status = false
	pf := models.FindPersonFilm(uid,movieId)
	if pf == nil {
		pf = new(models.PersonFilm)
		pf.MovieId = movieId
		pf.UserId = uid
		err := models.InsertPersonFilm(pf)
		if err != nil {
			resp.Msg = "个人对电影初始化失败！"
			return
		}
	}
	pf.Collect = collect
	err := models.UpdatePersonCollect(pf)
	if err != nil {
		resp.Msg = "操作失败，请稍后重试"
		return
	}
	collection,err := models.GetAvgCollect(movieId)
	if err != nil {
		resp.Msg = "获取电影评分最新状态失败"
		return
	}
	err = models.UpdateMoiveCollect(collection,movieId)
	if err != nil {
		resp.Msg = "更新电影状态失败"
	}
	resp.Status = true
	return
}
//个人对电影状态
func CollectStatus (uid string,movieId int)(resp MsgResponse){
	pf := models.FindPersonFilm(uid,movieId)
	if pf != nil {
		resp.Object = pf
	}else {
		resp.Object = models.PersonFilm{Collect:"false",Score:0}
	}
	resp.Status = true
	return
}
//评分
func ScoreFilm(uid string,movieId,score int)(resp MsgResponse){
	resp.Status = false
	pf := models.FindPersonFilm(uid,movieId)
	if pf == nil {
		pf = new(models.PersonFilm)
		pf.MovieId = movieId
		pf.UserId = uid
		err := models.InsertPersonFilm(pf)
		if err != nil {
			resp.Msg = "个人对电影初始化失败！"
			return
		}
	}
	pf.Score = score
	err := models.UpdatePersonScore(pf)
	if err != nil {
		resp.Msg = "操作失败，请稍后重试"
		return
	}
	socre ,err := models.GetAVGScore(movieId)
	if err != nil {
		resp.Msg = "获取电影评分最新状态失败"
		return
	}
	err = models.UpdateMovieScore(socre,movieId)
	if err != nil {
		resp.Msg = "更新电影状态失败"
		return
	}
	resp.Status = true
	return
}
//找到个人电影收藏
func FindLikeMovieByUid(userId string,page int)(resp MovieResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.UserCenterSize)
	m,err := models.FindLikeMovie(userId,pageIndex,util.UserCenterSize)
	if err != nil {
		resp.Msg = "查询出错了"
		return
	}
	count,err := models.CountLikeMoive(userId)
	if err != nil {
		resp.Msg = "总条数查询失败"
		return
	}
	pages := util.PageNum(count,util.UserCenterSize)
	ok :=pagination(&resp,pages,count,page,util.UserCenterSize)
	if !ok {
		return
	}
	GetMovieImgUrl(m)
	resp.Object = m
	resp.Status = true
	resp.Page = page
	return
}
//分页方法
func pagination(resp *MovieResp,pages,count,page,pageSize int)(bool){
	if page < pages {
		resp.Next = true
	}
	if count >pageSize && page != 1{
		resp.Pref = true
	}
	if count == 0 {
		resp.Msg = "没有数据"
		b := false
		return b
	}
	return true
}
