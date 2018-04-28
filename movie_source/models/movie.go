package models

import (
	"fmt"
	"zcm_tools/orm"
)

type Movie struct {
	Id          int     `description:"电影唯一id"`
	MovieName   string  `description:"电影名字"`
	Area        string  `description:"地区"`
	Language    string  `description:"语言"`
	Type        string  `description:"电影类型"`
	Director    string  `description:"导演"`
	ReleaseDate string  `description:"上映日期"`
	UpdateDate  string  `description:"更新日期"`
	Introuduce  string  `description:"电影简介"`
	Year        string  `description:"年份"`
	Plot        string  `description:"剧情"`
	Source      string  `description:"资源"`
	SourcePath  string  `description:"资源路径"`
	SourceFrom  int     `description:"资源来源"`
	Integral    int     `description:"积分"`
	MovieImg    string  `description:"电影海报"`
	Score       float64 `description:"电影评分"`
}

type Menu struct {
	Id   int    `description:"唯一编号"`
	Name string `description:"菜单名"`
}

//电影列表
func MovieList(pageIndex, pageSize int) (m []Movie, err error) {
	sql := `SELECT id,movie_name,release_date,source FROM movie LIMIT ?,? `
	_, err = orm.NewOrm().Raw(sql, pageIndex, pageSize).QueryRows(&m)
	return
}

//电影总数
func MovieListCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM movie`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//今日发布电影
func TodayMovie() (count int, err error) {
	sql := `SELECT COUNT(1) FROM movie WHERE update_date = CURDATE()`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//根据id找到电影
func FindMovieById(id int) (m Movie, err error) {
	sql := `SELECT * FROM movie WHERE id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&m)
	return
}

//全部区域
func FindAreaMenu() (a []Menu, err error) {
	sql := `SELECT * From area`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&a)
	return
}

//全部类型
func FindTypeMenu() (a []Menu, err error) {
	sql := `SELECT * From type`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&a)
	return
}

//全部语言
func FindLanguageMenu() (a []Menu, err error) {
	sql := `SELECT * From language`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&a)
	return
}

//修改电影
func ModifyMovie(m Movie, status int) (err error) {
	o := orm.NewOrm()
	o.Begin()
	defer func() {
		if err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}()
	sql := `UPDATE movie SET movie_name = ?,area = ?,language = ?,type = ?,director = ?,release_date = ?,introuduce = ?,year = ?,plot = ?
			WHERE id = ?`
	_, err = o.Raw(sql, m.MovieName, m.Area, m.Language, m.Type, m.Director, m.ReleaseDate, m.Introuduce, m.Year, m.Plot, m.Id).Exec()
	if status != 0 {
		sql = `UPDATE movie SET movie_img = ? WHERE id = ?`
		_, err = o.Raw(sql, m.MovieImg, m.Id).Exec()
	}
	fmt.Println(sql, m.Type)
	return
}

//更新资源
func UpdateMocieSource(m Movie) (err error) {
	sql := `UPDATE movie SET source_path = ?,source = 'exist',update_date = NOW() WHERE id = ?`
	_, err = orm.NewOrm().Raw(sql, m.SourcePath, m.Id).Exec()
	return
}

func InsertMovie(m Movie, status, videostatus int) (err error) {
	o := orm.NewOrm()
	o.Begin()
	defer func() {
		fmt.Println(err)
		if err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}()
	sql := `INSERT INTO movie (movie_name,area,language,type,director,release_date,introuduce,year,plot)
			VALUES
			(?,?,?,?,?,?,?,?,?)`
	_, err = o.Raw(sql, m.MovieName, m.Area, m.Language, m.Type, m.Director, m.ReleaseDate, m.Introuduce, m.Year, m.Plot).Exec()
	if err != nil {
		return
	}
	id := 0
	sql = `SELECT id FROM movie ORDER BY id DESC LIMIT 1`
	err = o.Raw(sql).QueryRow(&id)
	if err != nil {
		return
	}
	//初始化电影属性
	sql = `INSERT INTO movie_attribute (movie_id,year) VALUES (?,?)`
	_, err = o.Raw(sql, id, m.Year).Exec()
	if err != nil {
		return
	}
	//发布一条资源公告
	err = InsertOfficialMsgByMovie("电影资源更新！", "增加了<<"+m.MovieName+">>的资源小伙伴们快来啊！", o)
	if err != nil {
		return
	}
	if status != 0 {
		sql = `UPDATE movie SET movie_img = ? WHERE id = ?`
		_, err = o.Raw(sql, m.MovieImg, id).Exec()
		if err != nil {
			return
		}
	}
	if videostatus != 0 {
		sql = `UPDATE movie SET source_path = ?,source = 'exist',update_date = NOW() WHERE id = ?`
		_, err = o.Raw(sql, m.SourcePath, id).Exec()
		if err != nil {
			return
		}
	}

	fmt.Println(sql, m.SourcePath, id)
	return
}
