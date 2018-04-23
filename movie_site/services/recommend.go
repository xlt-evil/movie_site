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
