package services

import (
	"SB/movie_site/models"
	"SB/movie_site/util"
)
//随机推荐
func RandMovie()(resp MsgResponse){
	resp.Status = false
	length,err := models.CountMovie()
	if err != nil {
		resp.Msg ="无法查询电影总数"
		return
	}
	ids := util.RandomId(length)
	var movies []models.Movie
	for i,_:=range ids{
		m,_:=models.FindMovieByIds(ids[i])
		movies = append(movies,m)
	}
	resp.Object = movies
	resp.Status = true
	return
}
//热门电影推荐
func HotMovie()(resp MsgResponse){
	resp.Status = true
	m,err := models.HotMovie()
	if err != nil {
		resp.Msg = "Bug"
		return
	}
	resp.Object = m
	resp.Status = true
	return
}
//相关影评
func RelatedReview(id int)(resp MsgResponse){
	resp.Status = false
	f ,err := models.FindReviewByMovieId(id)
	if err != nil {
		resp.Msg = "bug"
		return
	}

	resp.Object = f
	resp.Status = true
	if len(f) == 0 {
		resp.Status = false
	}
 	return
}
//首页电影推荐
func IndexOffical()(resp MsgResponse){
	resp.Status = false
	m ,err := models.GetIndexOffical()
	if err != nil {
		resp.Msg = "bug"
		return
	}
	resp.Object = m
	resp.Status =  true
	return
}
//首页影评推荐
func IndexReview()(resp MsgResponse){
	resp.Status = false
	m ,err := models.FindHotReview()
	if err != nil {
		resp.Msg = "bug"
		return
	}
	resp.Object = m
	resp.Status =  true
	return
}
