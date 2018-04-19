package models

import (
	"fmt"
	"strconv"
	"zcm_tools/orm"
)

//用户表
type User struct {
	Id           int    `description:"注册顺序数"`
	Uid          string `description:"账户"`
	Birthday     string `description:"初始日期"`
	PersonSign   string `description:"个人签名"`
	PersonNotice string `description:"个人公告"`
	Name         string `description:"昵称"`
	Password     string `description:"密码"`
	Question     string `description:"密码问题"`
	Answer       string `descrition:"密码答案"`
	RegisterTime string `description:"注册时间"`
	ImgHead      string `description:"头像路径"`
	Integral     int    `description:"积分"`
}
type PersonFilm struct { //个人对电影表
	Id      int
	MovieId int
	UserId  string
	Score   int
	Collect string `description:"false ture 收藏电影"`
	//============================================
	ImgHead string `description:"user表的头像"`
}
//添加账户
func AddUser(u User) error {
	sql := `INSERT INTO user (uid,password,question,answer) VALUES(?,?,?,?)`
	o := orm.NewOrm()
	_, err := o.Raw(sql, u.Uid, u.Password, u.Question, u.Answer).Exec()
	return err
}
//查用户id是否存在
func CheckUid(uid string) (name string) {
	sql := `SELECT uid FROM user WHERE uid = ?`
	o := orm.NewOrm()
	o.Raw(sql, uid).QueryRow(&name)
	return
}
//判断用户账户密码
func CheckPwdAndUid(u User) (user *User) {
	sql := `SELECT * FROM user WHERE uid = ? AND password = ? `
	o := orm.NewOrm()
	o.Raw(sql, u.Uid, u.Password).QueryRow(&user)
	fmt.Print(user)
	return
}
//忘记密码
func ForgetPwd(u User) (password string) {
	sql := `SELECT password FROM user WHERE uid = ? AND question = ? AND answer = ?`
	o := orm.NewOrm()
	question, _ := strconv.Atoi(u.Question)
	o.Raw(sql, u.Uid, question, u.Answer).QueryRow(&password)
	return
}
//找到个人对电影
func FindPersonFilm(uid string,movieId int) (p *PersonFilm) {
	sql := `SELECT * FROM per_to_film
			WHERE user_id = ? AND movie_id = ? `
	o := orm.NewOrm()
	o.Raw(sql,uid,movieId).QueryRow(&p)
	return
}
//添加到个人对电影表
func InsertPersonFilm(p *PersonFilm)(err error) {
	sql := `INSERT INTO per_to_film (movie_id,user_id) VALUES (?,?)`
	o := orm.NewOrm()
	_,err = o.Raw(sql,p.MovieId,p.UserId).Exec()
	fmt.Println(sql,p)
	return
}
//改变收藏状态
func UpdatePersonCollect(p *PersonFilm)(err error){
	sql := `UPDATE per_to_film SET collect = ? WHERE user_id = ? AND movie_id = ? `
	o := orm.NewOrm()
	_,err =  o.Raw(sql,p.Collect,p.UserId,p.MovieId).Exec()
	return
}
//改变评分
func UpdatePersonScore(p *PersonFilm)(err error){
	sql := `UPDATE per_to_film SET score = ? WHERE user_id = ? AND movie_id = ? `
	o := orm.NewOrm()
	_,err =  o.Raw(sql,p.Score,p.UserId,p.MovieId).Exec()
	return
}
//找个人收藏电影
func FindLikeMovie(userId string,pageIndex,pageSize int)(m []Movie,err error){
	sql := `SELECT
			m.id,m.movie_name,m.movie_img
			FROM
				per_to_film pf
			LEFT JOIN movie m ON pf.movie_id = m.id
			WHERE
			pf.user_id = ?
			LIMIT ?,?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,userId,pageIndex,pageSize).QueryRows(&m)
	fmt.Println(sql,userId,pageIndex,pageSize,m)
	return
}
//计算个人收藏电影总数
func CountLikeMoive(userId string)(count int,err error){
	sql := `SELECT
			COUNT(1)
			FROM
				per_to_film pf
			LEFT JOIN movie m ON pf.movie_id = m.id
			WHERE
			pf.user_id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,userId).QueryRow(&count)
	return
}



type MyReview struct {
	Id int
	MovieName string
	Date string
	Title string
}
//个人影评
func GetMyReview(uid string,pageIndex,pageSize int)(r []MyReview,err error){
	sql := `SELECT m.movie_name,f.title,f.id,f.date
			FROM film_review f
			LEFT JOIN movie m on f.movie_id = m.id
			WHERE f.user_id = ?
			ORDER BY f.date DESC
			LIMIT ?,?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,uid,pageIndex,pageSize).QueryRows(&r)
	return
}

//个人影评
func GetMyReviewCount(uid string)( count int,err error){
	sql := `SELECT COUNT(1)
			FROM film_review f
			WHERE f.user_id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,uid).QueryRow(&count)
	return
}