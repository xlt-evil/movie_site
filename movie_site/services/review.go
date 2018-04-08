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

