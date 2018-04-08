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
//电影短评表发布
func MovieTalk(mt *models.MovieTalk)(resp MsgResponse){
	resp.Status = false
	f := models.FindPersonFilm(mt.UserId,mt.MovieId)
	if f == nil {
		f = new(models.PersonFilm)
		f.MovieId = mt.MovieId
		f.UserId = mt.UserId
		models.InsertPersonFilm(f)
	}
	err := models.InsertNewMovieTalk(mt)
	if err != nil {
		resp.Msg = "增加内容失败"
		return
	}
	resp.Status = true
	return
}
//影评发布
func ReleaseFilmCritic(fr *models.FilmReview)(resp MsgResponse){
	resp.Status = false
	err := models.InsertFilmCritic(fr)
	if err != nil {
		resp.Msg = "发布影评失败"
		return
	}
	id,err := models.FindNewFilmReview(fr.MovieId,fr.UserId)
	if err!=nil {
		resp.Msg = "发布影评失败了"
		return
	}
	err = models.InsertFilmReviewAttribute(id,fr.MovieId)
	if err != nil {
		resp.Msg = "初始化影评属性失败"
		return
	}
	resp.Status = true
	return
}
//影评喜好选择
func ReviewFeel(fr_id int,user_id string,feel int)(resp MsgResponse){
	resp.Status = false
	err := models.InsertPerToReview(fr_id,user_id)
	if err != nil {
		resp.Msg = "个人对影评喜好初始化失败"
		return
	}
	err = models.UpdatePerToReview(fr_id,user_id,feel)
	if err != nil {
		resp.Msg = "个人对影评喜欢改变失败"
		return
	}
	hate,err :=models.CountFeelIsHate(fr_id)
	like,err :=models.CountFeelIsLike(fr_id)
	if err != nil {
		resp.Msg = "获取状态失败"
		return
	}
	err = models.UpdateFilmAttribute(fr_id,hate,like)
	if err != nil {
		resp.Msg = "更新状态失败"
		return
	}
	resp.Status = true
	return

}