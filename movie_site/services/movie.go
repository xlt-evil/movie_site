package services

import (
	"SB/movie_site/models"
	"SB/movie_site/util"
)

//电影的基本服务

//找到电影
func FindMovie(id int)(resp MsgResponse){
	m,err := models.FindMovieById(id)
	if err != nil {
		resp.Status = false
		resp.Msg = "资源出现问题，请稍后重试"
		return
	}
	//增加电影热度
	err = models.UpPopularity(id)
	if err != nil {
		resp.Msg = "电影热度增加失败"
		return
	}
	resp.Status = true
	resp.Object = m
	return
}

//电影属性
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

//热门电影首次加载
func FindMovieListByIndex(key []string)(resp MsgResponse){
	m,err := models.FindMovieListByIndex(key)
	if err != nil {
		resp.Status = false
		resp.Msg = "服务器出现问题，请稍后重试"
		return
	}
	resp.Status = true
	GetMovieImgUrl(m)
	resp.Object = m
	return
}

//热门电影异步加载
func FindMovieListByIndexAjax(condition string,paras ...interface{})(resp MsgResponse){
	m,err := models.FindMovieListByIndexByAjax(condition,paras)
	if err != nil {
		resp.Status = false
		resp.Msg = "热门电影加载失败！"
		return
	}
	resp.Status = true
	GetMovieImgUrl(m)
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

type MovieResp struct {
	MsgResponse
	Pref bool //是否有上一页
	Next bool //是否有下一页
	Page int  //当前页
}

//电影库异步加载
func MovieLibraryAjax(area,types,page int ,year,condition,orderby string)(resp MovieResp){
	var err error
	var m []models.Movie
	paras := []interface{}{}
	if area != 0{
		a,err := models.GetAreaName(area)
		resp.Status = false
		if err != nil {
			resp.Msg = "参数转换失败，请稍后重试area"
			return
		}
		paras = append(paras,a.Name)
	}
	if types != 0{
		t,err := models.GetTypeName(types)
		if err != nil {
			resp.Msg = "参数转换失败，请稍后重试types"
			return
		}
		paras = append(paras,t.Name)

	}
	pageIndex := util.PageIndex(page,util.MoviePageSize)
	if year != "全部年代" && year != "2007以前" {
		paras = append(paras,year)
	}

	if pageIndex == 0{
		m,err = models.MovieLibrary(condition,orderby,pageIndex,util.MoviePageSize,paras)
	}else{
		m,err = models.MovieLibrary(condition,orderby,pageIndex,util.MoviePageSize,paras)
	}
	if err != nil {
		resp.Msg = "电影列表获取失败"
		return
	}
	num ,err:= models.CountMovieLibrary(condition,paras)
	if err != nil {
		resp.Msg = "电影总数获取失败"
		return
	}
    pages := util.PageNum(num,util.MoviePageSize)
	if page < pages {
		resp.Next = true
	}
	if num > 12 && page != 1 {
		resp.Pref = true
	}
	resp.Page = page
	GetMovieImgUrl(m)
	resp.Status = true
	resp.Object = m
	if num == 0{
		resp.Status = false
	}
	return
}

//首次到电影页面
func MovieLibrary()(resp MsgResponse){
	resp.Status = false
	m ,err := models.MovieLibraryPage()
	if err != nil {
		resp.Msg = "电影列表获取失败"
		return
	}
	GetMovieImgUrl(m)
	resp.Object = m
	return
}

//查看电影名字
func FindMovieName(id int )(name string,movieId int){
	name,_ = models.FindMovieName(id)
	movieId,_ = models.FindMovieId(id)
	return
}





