package service

import (
	"movie_source/util"
	"movie_source/models"
	"time"
)
//电影详情
func MovieList(page int )(resp PagemsgResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.PageSize4)
	m ,err := models.MovieList(pageIndex,util.PageSize4)
	if err != nil {
		resp.Msg = "查询建议列表时出现bug了"
		return
	}
	count,err := models.MovieListCount()
	if err != nil {
		resp.Msg = "查询总建议列表条数时出现bug了"
		return
	}
	pages := util.PageNum(count,util.PageSize4)
	pagination(&resp,pages,count,page,util.PageSize4)
	resp.Status = true
	resp.Page = page
	resp.Object = m
	return
}

//今总电影数
func MovieCount()(today,count int){
	count,_= models.MovieListCount()
	today,_= models.TodayMovie()
	return
}
//电影具体
func GetMovieDetail(id int)(resp MsgResp){
	resp.Status = false
	m,err := models.FindMovieById(id)
	if err != nil {
		resp.Msg = "查看电影详情失败"
		return
	}
	resp.Object = m
	resp.Status = true
	return
}
//菜单渲染
func Menu()(area ,types,language []models.Menu){
	area ,_= models.FindAreaMenu()
	types ,_ = models.FindTypeMenu()
	language,_ = models.FindLanguageMenu()
	return
}
//修改电影
func ModifyMovie(movieId int,moviename,director,area,types,language,releasedate,content,short,img string,status int)(resp MsgResp){
	y ,_ := time.Parse("2006-01-02",releasedate)
	year := y.Format("2006-01-02")
	err := models.ModifyMovie(models.Movie{
		Id:movieId,
		MovieName:moviename,
		Director:director,
		Area:area,
		Type:types,
		Language:language,
		Introuduce:short,
		Plot:content,
		Year:year,
		ReleaseDate:releasedate,
		MovieImg:img,
		},status)
	if err != nil {
		resp.Msg = "修改失败"+err.Error()
		return
	}
	resp.Status = true
	return
}

func UpdateMovieSource(sourcepath string,id int)(resp MsgResp){
	resp.Status = false
	err := models.UpdateMocieSource(models.Movie{
		Id:id,
		SourcePath:sourcepath,
	})
	if err != nil {
		resp.Msg = "上传失败"
		return
	}
	resp.Status = true
	return
}

func ReleaseMovie(moviename,director,area,types,language,releasedate,content,short,img string,status,videostatus int,sourcepath string)(resp MsgResp){
	y ,_ := time.Parse("2006-01-02",releasedate)
	year := y.Format("2006-01-02")
	resp.Status = true
	err := models.InsertMovie(models.Movie{
		MovieName:moviename,
		Director:director,
		Area:area,
		Type:types,
		Language:language,
		ReleaseDate:releasedate,
		Introuduce:short,
		Plot:content,
		Year:year,
		MovieImg:img,
		SourcePath:sourcepath,
	},status,videostatus)
	if err != nil {
		resp.Msg = "添加电影失败！"
		return
	}
	return
}

