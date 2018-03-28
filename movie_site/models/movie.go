package models

import (
	"movie_site/util"
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
}

type MoiveAttribute struct {
	Id         int     `description:"id"`
	MovieId    int     `description:"电影唯一id"`
	Popularity int     `description:"热度"`
	CollectNum int     `description:"收藏人数"`
	Score      float64 `description:"评分"`
	Year       string  `description:"年份"`
}

func FindMovieById(id int) (m Movie, err error) {
	sql := `SELECT * FROM movie WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&m)
	m.MovieImg = util.Movie_img + m.MovieImg
	m.SourcePath = util.Source_path+m.SourcePath
	return
}

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
	fmt.Print(sql,key[0],key[1],err,movies)
	return
}
