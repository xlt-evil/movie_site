package models

import (
	"SB/movie_site/util"
	"zcm_tools/orm"
	"fmt"
)

type Movie struct {
	Id          int    `description:"电影唯一id"`
	MovieName   string `description:"电影名字"`
	Area        string `description:"地区"`
	Language    string `description:"语言"`
	Type        string `description:"电影类型"`
	Director    string `description:"导演"`
	ReleaseDate string `description:"上映日期"`
	UpdateDate  string `description:"更新日期"`
	Introuduce  string `description:"电影简介"`
	Year        string `description:"年份"`
	Plot        string `description:"剧情"`
	Source      string `description:"资源"`
	SourcePath  string `description:"资源路径"`
	SourceFrom  int    `description:"资源来源"`
	Integral    int    `description:"积分"`
	MovieImg    string `description:"电影海报"`
	Score       float64    `description:"电影评分"`
}

type MoiveAttribute struct {
	Id         int     `description:"id"`
	MovieId    int     `description:"电影唯一id"`
	Popularity int     `description:"热度"`
	CollectNum int     `description:"收藏人数"`
	Score      float64 `description:"评分"`
	Year       string  `description:"年份"`
}

//地区
type Area struct {
	Id int
	Name string
}

//类型
type Type struct {
	Id int
	Name string
}
//根据id找到电影
func FindMovieById(id int) (m Movie, err error) {
	sql := `SELECT * FROM movie WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&m)
	m.MovieImg = util.Movie_img + m.MovieImg
	m.SourcePath = util.Source_path+m.SourcePath
	return
}
//找到电影的属性
func FindMovieAttribute(id int) (attribute MoiveAttribute, err error) {
	sql := `SELECT * FROM movie_attribute WHERE movie_id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&attribute)
	return
}
//首页热门电影 因为要加载到缓存所以单独写
func FindMovieListByIndex(key []string)(movies []Movie,err error){
	sql := `SELECT
				m.*
			FROM
				(
					SELECT
						*
					FROM
						movie
					WHERE
						area = ?
					OR
						area = ?
				) m
			LEFT JOIN movie_attribute ma ON ma.movie_id = m.id
			ORDER BY
				ma.popularity DESC,
				ma.id ASC
			LIMIT 0,8`
	o := orm.NewOrm()
	_, err = o.Raw(sql,key[0],key[1]).QueryRows(&movies)
	return
}
//热门电影异步加载
func FindMovieListByIndexByAjax(condition string,paras ...interface{})(movies []Movie,err error){
	sql := `SELECT
				m.*
			FROM
				(
					SELECT
						*
					FROM
						movie
					 `
			if condition != ""{
				sql += condition
			}
			sql+=` ) m
			LEFT JOIN movie_attribute ma ON ma.movie_id = m.id
			ORDER BY
				ma.popularity DESC,
				ma.id ASC
			LIMIT 0,8`
	o := orm.NewOrm()
	_, err = o.Raw(sql,paras).QueryRows(&movies)
	fmt.Println(sql,paras,movies)
	return
}
//电影库异步加载
func MovieLibrary(condition,orderby string , pageIndex,pageSize int,paras ...interface{})(m []Movie,err error){
	sql :=  `SELECT m.*,ma.score FROM (
					SELECT * FROM movie
					WHERE 1=1 `
	if condition != ""{
		sql += condition
	}
	sql+=` ) m LEFT JOIN movie_attribute ma ON ma.movie_id = m.id `
	sql+= orderby
	sql+=" LIMIT ?,?"
	o := orm.NewOrm()
	_,err = o.Raw(sql,paras,pageIndex,pageSize).QueryRows(&m)
	fmt.Println(sql,paras,pageIndex,pageSize)
	return
}
//得到地区名
func GetAreaName(id int)(a Area,err error){
	sql := `SELECT * FROM area WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,id).QueryRow(&a)
	fmt.Println(sql,id,a)
	return
}
//得到类型名
func GetTypeName(id int)(a Type,err error){
	sql := `SELECT * FROM type WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,id).QueryRow(&a)
	return
}
//电影库首次加载页面
func MovieLibraryPage()(m []Movie,err error){
	sql := `SELECT m.*,ma.score FROM
				(SELECT * FROM movie ) m
				LEFT JOIN movie_attribute ma on ma.movie_id = m.id
				ORDER BY ma.popularity DESC,
                ma.id ASC
				LIMIT 0 ,12`
	o := orm.NewOrm()
	_,err = o.Raw(sql).QueryRows(&m)
	return
}
//按条件电影总数
func CountMovieLibrary(condition string ,paras ...interface{})(count int,err error){
	sql :=  `SELECT COUNT(1) FROM (
					SELECT * FROM movie
					WHERE 1=1 `
	if condition != ""{
		sql += condition
	}
	sql+=` ) m LEFT JOIN movie_attribute ma ON ma.movie_id = m.id `
	o := orm.NewOrm()
	err = o.Raw(sql,paras).QueryRow(&count)
	return
}
//更新电影熟悉分数表的状态
func UpdateMovieScore(score ,movieId int)(err error){
	sql := `UPDATE movie_attribute SET score = ? WHERE movie_id = ? `
	o := orm.NewOrm()
	_,err = o.Raw(sql,score,movieId).Exec()
	return
}
//获取摸个电影的平均值
func GetAVGScore(movieId int)(socre int,err error) {
	sql := `SELECT ROUND(AVG(score)) socre from per_to_film WHERE movie_id = ? `
	o := orm.NewOrm()
	err = o.Raw(sql, movieId).QueryRow(&socre)
	return
}
//获取某个电影的收藏人数
func GetAvgCollect(movieId int)(collect int ,err error){
	sql := `SELECT count(1) collect from per_to_film WHERE movie_id = ? AND collect <> 2`
	o := orm.NewOrm()
	err = o.Raw(sql, movieId).QueryRow(&collect)
	return
}
//更新某个电影的收藏人数
func UpdateMoiveCollect(collect ,movieId int)(err error){
	sql := `UPDATE movie_attribute SET collect_num = ? WHERE movie_id = ? `
	o := orm.NewOrm()
	_,err = o.Raw(sql,collect,movieId).Exec()
	return
}
//查看电影名
func FindMovieName(movieId int)(name string,err error){
	sql := `SELECT movie_name name FROM movie WHERE id = ? `
	o := orm.NewOrm()
	err = o.Raw(sql,movieId).QueryRow(&name)
	return
}
//查看电影id
func FindMovieId(movieId int)(id int,err error){
	sql := `SELECT id FROM movie WHERE id = ? `
	o := orm.NewOrm()
	err = o.Raw(sql,movieId).QueryRow(&id)
	return
}
//查看电影总数
func CountMovie()(count int,err error){
	sql := `SELECT COUNT(1)	FROM movie`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}
//随机推荐
func FindMovieByIds(id int)(m Movie,err error){
	sql := `SELECT movie_name,ma.score,m.id FROM movie m
			LEFT JOIN movie_attribute ma ON ma.movie_id = m.id
			WHERE m.id = ?`
	err = orm.NewOrm().Raw(sql,id).QueryRow(&m)
	return
}
//热门电影
func HotMovie()(m []Movie,err error){
	sql := `SELECT movie_name,ma.score,m.id FROM movie m
			LEFT JOIN movie_attribute ma ON ma.movie_id = m.id
			ORDER BY ma.popularity DESC
			LIMIT 0,8`
	_,err = orm.NewOrm().Raw(sql).QueryRows(&m)
	return

}
//真假电影人气
func UpPopularity(id int)(err error){
	sql := `UPDATE movie_attribute SET popularity = popularity+1
			WHERE movie_id = ?`
	_,err = orm.NewOrm().Raw(sql,id).Exec()
	return
}