package services

import (
	"SB/movie_site/util"
	"SB/movie_site/models"
	"fmt"
)
//评论
//电影短评异步加载
func MovieTalkAjax(movieId ,page int)(resp MovieResp){
	resp.Status = false
	pageIndex  := util.PageIndex(page,util.MovieTalkPageSize)
	m,err := models.FindMovieTalk(movieId,pageIndex,util.MovieTalkPageSize)
	if err != nil {
		resp.Msg = "查询短评失败"
		return
	}
	num ,err:= models.CountMoiveTalk(movieId)
	if err != nil {
		resp.Msg = "获取全部短评条数失败!"
		return
	}
	pages := util.PageNum(num,util.MovieTalkPageSize)
	if page < pages {
		resp.Next = true
	}
	if num > 8 && page != 1 {
		resp.Pref = true
	}
	for i,_ :=range m {
		m[i].ImgHead = util.Person_img+m[i].ImgHead
	}

	if num == 0{
		resp.Msg = "暂无短评哦！快来占沙发呀"
		return
	}
	resp.Page = page
	resp.Status = true
	resp.Object = m
	return
}
//电影评论列表加载
func FindFilmReview(page int)(resp MovieResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.MovieTalkPageSize)
	m ,err:=models.FindFilmReviewList(pageIndex,util.MovieTalkPageSize)
	if err != nil {
		resp.Msg = "查询影评列表失败"
		return
	}
	count ,err := models.CountFilmReview()
	if err != nil {
		resp.Msg = "查询总条数失败"
		return
	}
	pages := util.PageNum(count,util.MovieTalkPageSize)

	if page < pages {
		resp.Next = true
	}
	if count > 8 && page != 1 {
		resp.Pref = true
	}
	for i,_ :=range m {
		m[i].ImgHead = util.Person_img+m[i].ImgHead
		m[i].MovieImg = util.Movie_img+m[i].MovieImg
	}

	if count == 0{
		resp.Msg = "暂无短评哦！快来占沙发呀"
		return
	}
	resp.Object = m
	resp.Status = true
	resp.Page = page
	fmt.Println(resp.Pref,resp.Next)
	return
}
//具体影评页面加载
func TalkDetail(id int)(resp MsgResponse){
	resp.Status = false
	talk,err := models.GetFilmReviewById(id)
	if err != nil {
		resp.Msg = "读取该影评内容失败"
		return
	}
	talk.ImgHead = util.Person_img+talk.ImgHead
	resp.Object = talk
	resp.Status = true
	fmt.Println(talk.Name)
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
	p,_ :=models.FindPerToReview(fr_id,user_id)
	if p == nil {
		err := models.InsertPerToReview(fr_id,user_id)
		if err != nil {
			resp.Msg = "个人对影评喜好初始化失败"
			return
		}
	}else{
		resp.Msg = "不能重复点赞"
		return
	}
	err := models.UpdatePerToReview(fr_id,user_id,feel)
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
	err = models.UpdateFilmAttribute(fr_id,like,hate)
	fmt.Println(err)
	if err != nil {
		resp.Msg = "更新状态失败"
		return
	}
	resp.Status = true
	fmt.Println("123")
	return

}
//影评短评发布
func ReviewsRelease(fr_id,movieId int ,userId ,content string)(resp MsgResponse){
	resp.Status = false
	err:=models.InsertFrTalk(fr_id,movieId,userId,content)
	if err != nil {
		resp.Msg = "发布失败"
		return
	}
	count,err := models.CountFrTalk(fr_id)
	err = models.UpdateReviewNum(fr_id,count)
	if err != nil {
		resp.Msg = "增加该影评总条数失败"
		return
	}
	resp.Status = true
	return
}
//影评列表加载
func LoadReleaseList(page int,fr_id int)(resp MovieResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.MovieTalkPageSize)
	m ,err := models.FindFrTalkByFrid(fr_id,pageIndex,util.MovieTalkPageSize)
	if err != nil {
		resp.Msg = "查询影评评论列表失败"
		return
	}
	count ,err := models.CountFrTalk(fr_id)
	if err != nil {
		resp.Msg = "查询总条数失败"
		return
	}
	pages := util.PageNum(count,util.MovieTalkPageSize)

	if page <pages {
		resp.Next = true
	}
	if count > 8 && page != 1 {
		resp.Pref = true
	}
	for i,_:=range m {
		m[i].ImgHead = util.Person_img+m[i].ImgHead
	}
	if count == 0 {
		resp.Msg =  "暂无短评哦！快来占沙发呀"
		return
	}
	resp.Object = m
	resp.Status = true
	resp.Page = page
	return
}
