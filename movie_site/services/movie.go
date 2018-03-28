package services

import (
	"movie_site/models"
	"movie_site/util"
	"fmt"
)

//电影

func FindMovie(id int)(resp MsgResponse){
	m,err := models.FindMovieById(id)
	if err != nil {
		resp.Status = false
		resp.Msg = "资源出现问题，请稍后重试"
		return
	}
	resp.Status = true
	resp.Object = m
	return
}

func FindMovieAttribute(id int)(resp MsgResponse){
	m,err := models.FindMovieAttribute(id)
	if err != nil {
		resp.Status = false
		resp.Msg = "资源出现问题，请稍后重试"
		return
	}
	resp.Status = true
	resp.Object = m
	return
}

func FindMovieListByIndex(key []string)(resp MsgResponse){
	m,err := models.FindMovieListByIndex(key)
	if err != nil {
		resp.Status = false
		resp.Msg = "服务器出现问题，请稍后重试"
		return
	}
	resp.Status = true
	GetMovieImgUrl(m)
	fmt.Print(m[0].MovieImg)
	resp.Object = m
	return
}

//常用方法
func GetMovieImgUrl(m []models.Movie)([]models.Movie){
	for i,_ := range m {
		m[i].MovieImg = util.Movie_img+m[i].MovieImg
	}
	return m
}


