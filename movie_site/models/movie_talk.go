package models

import (
	"fmt"
	"time"
	"zcm_tools/orm"
)

//电影短评
type MovieTalk struct {
	Id      int
	MovieId int       `description:"电影id"`
	UserId  string    `description:"用户id"`
	Date    time.Time `description:"时间"`
	Content string    `description:"短评内容"`
	Score   float64
	ImgHead string
	Name    string
}

//影评表
type FilmReview struct {
	Id      int
	MovieId int       `description:"电影id"`
	UserId  string    `description:"用户id"`
	Title   string    `description:"标题"`
	Content string    `description:"内容"`
	Date    time.Time `description:"日期"`
}

//影评列表 用于影评页面场所
type FilmReviewList struct {
	Id        int
	MovieId   int
	UserId    int
	MovieImg  string `description:"电影海报"`
	ImgHead   string `description:"用户头像"`
	MovieName string `description:"电影名"`
	Name      string `description:"用户名"`
	Score     int    `description:"评分"`
	Date      string `description:"发布时间"`
	Content   string `description:"内容"`
	Title     string `description:"影评标题"`
	Like      int    `description:"喜欢"`
	Hate      int    `description:"不喜欢"`
	TalkNum   int    `description:"评论人数"`
}

//影评属性表
type FilmReviewAttribute struct {
	Id      int
	FrId    int
	MovieId int
	Like    int `description:"喜欢人数"`
	Hate    int `dewcription:"讨厌人数"`
	TalkNum int `description:"谈论次数"`
}

//个人对影评表
type PerToReview struct {
	Id     int
	FrId   int
	UserId int
	Status string `description:"点击状态"`
	Feel   string `description:"喜好"`
	Date   time.Time
}

//影评评论表
type FrTalk struct {
	Id      int
	FrId    int
	MovieId int
	UserId  int
	Content string
	Date    time.Time
}

//影评影评页面展示
type FrTalkShow struct {
	Name    string
	Score   int
	Date    time.Time
	ImgHead string
	Content string
}

//个人中心影评展示
type UserCenterReview struct {
	Id      int    `description:"影评id"`
	Content string `description:"内容"`
	Date    string `description:"日期"`
	Name    string `description:"昵称"`
	ImgHead string `description:"头像"`
	Title   string `description:"标题"`
}

//找到该电影的所有短评
func FindMovieTalk(movieId, pageIndex, PageSize int) (mt []MovieTalk, err error) {
	sql := `SELECT
				u.img_head,p.score,mc.content,u.name
			FROM
				per_to_film p
			LEFT JOIN USER u ON p.user_id = u.uid
			LEFT JOIN movie_commentary mc ON p.movie_id = mc.movie_id AND p.user_id = mc.user_id
			where mc.movie_id = ?
			ORDER BY mc.date desc
			LIMIT ?,? `
	o := orm.NewOrm()
	_, err = o.Raw(sql, movieId, pageIndex, PageSize).QueryRows(&mt)
	fmt.Println("秀起来", sql, movieId, pageIndex, PageSize, mt)
	return
}

//插入一条短评
func InsertNewMovieTalk(mt *MovieTalk) (err error) {
	sql := `INSERT INTO movie_commentary (movie_id,user_id,content,date) VALUES (?,?,?,NOW())`
	o := orm.NewOrm()
	_, err = o.Raw(sql, mt.MovieId, mt.UserId, mt.Content).Exec()
	return
}

//该电影的评论总数
func CountMoiveTalk(movieId int) (count int, err error) {
	sql := `SELECT COUNT(1) FROM movie_commentary WHERE movie_id = ? `
	o := orm.NewOrm()
	err = o.Raw(sql, movieId).QueryRow(&count)
	return
}

//发布影评
func InsertFilmCritic(fr *FilmReview) (err error) {
	sql := `INSERT INTO film_review (movie_id,user_id,title,content,date) values(?,?,?,?,NOW())`
	o := orm.NewOrm()
	_, err = o.Raw(sql, fr.MovieId, fr.UserId, fr.Title, fr.Content).Exec()
	return
}

//电影影评列表
func FindFilmReviewList(pageIndex, pageSize int) (fr []FilmReviewList, err error) {
	sql := `SELECT
				fr.id,m.movie_img ,u.img_head,m.movie_name,u.NAME name,pf.score,fr.date,fr.content,fr.title,fr.movie_id,fr.user_id,fa.like,fa.hate,fa.talk_num
			FROM
				film_review fr
				LEFT JOIN movie m ON m.id = fr.movie_id
				LEFT JOIN USER u ON u.uid = fr.user_id
				LEFT JOIN per_to_film pf ON pf.movie_id = fr.movie_id
				AND pf.user_id = fr.user_id
				LEFT JOIN fr_attribute fa ON fa.fr_id = fr.id
				ORDER BY
				date DESC,
				fr.id ASC
				LIMIT ?,? `
	o := orm.NewOrm()
	_, err = o.Raw(sql, pageIndex, pageSize).QueryRows(&fr)
	fmt.Println(fr)
	return
}

//计算所有影评条数
func CountFilmReview() (count int, err error) {
	sql := `SELECT
				COUNT(1)
			FROM
				film_review fr
				LEFT JOIN movie m ON m.id = fr.movie_id
				LEFT JOIN USER u ON u.uid = fr.user_id
				LEFT JOIN per_to_film pf ON pf.movie_id = fr.movie_id
				AND pf.user_id = fr.user_id`
	o := orm.NewOrm()
	err = o.Raw(sql).QueryRow(&count)
	return
}

//找到该影评内容
func GetFilmReviewById(id int) (fr FilmReviewList, err error) {
	sql := `SELECT
				fr.id,m.movie_img ,u.img_head,m.movie_name,u.NAME name,pf.score,fr.date,fr.content,fr.title,fr.movie_id,fr.user_id,fa.like,fa.hate,fa.talk_num
			FROM
				film_review fr
				LEFT JOIN movie m ON m.id = fr.movie_id
				LEFT JOIN USER u ON u.uid = fr.user_id
				LEFT JOIN per_to_film pf ON pf.movie_id = fr.movie_id
				AND pf.user_id = fr.user_id
				LEFT JOIN fr_attribute fa ON fa.fr_id = fr.id
				WHERE fr.id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, id).QueryRow(&fr)
	return
}

//插入一个属性表
func InsertFilmReviewAttribute(id, movieId int) (err error) {
	sql := `INSERT INTO fr_attribute (fr_id,movie_id) values(?,?)`
	o := orm.NewOrm()
	_, err = o.Raw(sql, id, movieId).Exec()
	return
}

//找到最新的
func FindNewFilmReview(MovieId int, userId string) (id int, err error) {
	sql := `SELECT id from film_review WHERE movie_id = ? AND user_id = ? ORDER BY date desc LIMIT 0,1`
	o := orm.NewOrm()
	err = o.Raw(sql, MovieId, userId).QueryRow(&id)
	return
}

//插入个人对影评表
func InsertPerToReview(frid int, user_id string) error {
	sql := `INSERT INTO per_to_review (fr_id,user_id,date,status) VALUES (?,?,NOW(),'false')`
	o := orm.NewOrm()
	_, err := o.Raw(sql, frid, user_id).Exec()
	return err
}

//改变个人对影评表
func UpdatePerToReview(frid int, user_id string, feel int) error {
	sql := `UPDATE per_to_review SET feel = ?, status = 'true' WHERE fr_id = ? AND user_id = ? AND status = 'false'`
	o := orm.NewOrm()
	_, err := o.Raw(sql, feel, frid, user_id).Exec()
	return err
}

//统计影评讨厌人数表
func CountFeelIsHate(fr_id int) (hate int, err error) {
	sql := `SELECT COUNT(feel) FROM per_to_review where feel = 'hate' AND fr_id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, fr_id).QueryRow(&hate)
	return
}

//统计影评喜欢人数表
func CountFeelIsLike(fr_id int) (like int, err error) {
	sql := `SELECT COUNT(feel) FROM per_to_review where feel = 'like' AND fr_id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, fr_id).QueryRow(&like)
	return
}

//更新影评属性表
func UpdateFilmAttribute(fr_id, like, hate int) error {
	sql := `UPDATE fr_attribute fr SET fr.like = ?,fr.hate = ?
			WHERE fr_id = ?`
	o := orm.NewOrm()
	_, err := o.Raw(sql, like, hate, fr_id).Exec()
	return err
}

//找到个人对电影
func FindPerToReview(fr_id int, user_id string) (p *PerToReview, err error) {
	sql := `SELECT * FROM per_to_review WHERE fr_id = ? AND user_id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, fr_id, user_id).QueryRow(&p)
	return
}

//插入个人对影评评论
func InsertFrTalk(fr_id, movie_id int, user_id, content string) (err error) {
	sql := `INSERT INTO fr_talk (fr_id,movie_id,user_id,content,date) VALUES (?,?,?,?,NOW())`
	o := orm.NewOrm()
	_, err = o.Raw(sql, fr_id, movie_id, user_id, content).Exec()
	fmt.Println(sql, fr_id, movie_id, user_id, content)
	return
}

//找到该影评评论的全部条数
func CountFrTalk(fr_id int) (count int, err error) {
	sql := `SELECT COUNT(1) count FROM fr_talk WHERE fr_id = ?`
	o := orm.NewOrm()
	err = o.Raw(sql, fr_id).QueryRow(&count)
	return
}

//影评评论分页
func FindFrTalkByFrid(fr_id, pageIndex, pageSize int) (list []FrTalkShow, err error) {
	sql := `SELECT
 			u.name,pf.score,fa.date,u.img_head,fa.content
			FROM fr_talk fa
			LEFT JOIN user u on fa.user_id = u.uid
			LEFT JOIN per_to_film pf on pf.user_id = fa.user_id AND pf.movie_id  = fa.movie_id
			WHERE fa.fr_id = ?
			ORDER by fa.date desc
			LIMIT ?,?`
	o := orm.NewOrm()
	_, err = o.Raw(sql, fr_id, pageIndex, pageSize).QueryRows(&list)
	return
}

//增加到影评评论属性
func UpdateReviewNum(fr_id int, count int) (err error) {
	sql := `UPDATE fr_attribute fr SET fr.talk_num = ?
			WHERE fr_id = ?`
	o := orm.NewOrm()
	_, err = o.Raw(sql, count, fr_id).Exec()
	return
}

//找到该用户的所有评论
func FindReviewByUid(uid string, pageIndex, pageSize int)(r []UserCenterReview,err error){
	sql := `SELECT fr.id,ft.content,ft.date,u.name,u.img_head,fr.title FROM film_review fr
			RIGHT  JOIN fr_talk ft ON fr.id = ft.fr_id
			LEFT JOIN user u ON u.uid = ft.user_id
			WHERE fr.user_id = ?
			ORDER BY date DESC
			LIMIT ?,?`
	_,err = orm.NewOrm().Raw(sql, uid, pageIndex, pageSize).QueryRows(&r)
	return
}

//找到该用户的所有评论总条数
func FindReviewByUidCount(uid string)(count int,err error){
	sql := `SELECT count(1) FROM film_review fr
			RIGHT  JOIN fr_talk ft ON fr.id = ft.fr_id
			LEFT JOIN user u ON u.uid = ft.user_id
			WHERE fr.user_id = ?`
	err = orm.NewOrm().Raw(sql,uid).QueryRow(&count)
	return
}

//找打该电影的相关影评
func FindReviewByMovieId(id int )(f []FilmReview,err error){
	sql := `SELECT id,title FROM film_review
			WHERE movie_id = ?
			LIMIT 0,5`
	_,err = orm.NewOrm().Raw(sql,id).QueryRows(&f)
	return
}
//找到热门影评
func FindHotReview()(f []FilmReview,err error){
	sql := `SELECT fr.id,title FROM film_review fr
			LEFT JOIN fr_attribute fa ON fr.id = fa.fr_id
			ORDER BY fa.like DESC
			LIMIT 0,7`
	_,err = orm.NewOrm().Raw(sql).QueryRows(&f)
	return
}

//电影影评列表
func FindFilmReviewList2(pageIndex, pageSize int) (fr []FilmReviewList, err error) {
	sql := `SELECT
				fr.id,m.movie_img ,u.img_head,m.movie_name,u.NAME name,pf.score,fr.date,fr.content,fr.title,fr.movie_id,fr.user_id,fa.like,fa.hate,fa.talk_num
			FROM
				film_review fr
				LEFT JOIN movie m ON m.id = fr.movie_id
				LEFT JOIN USER u ON u.uid = fr.user_id
				LEFT JOIN per_to_film pf ON pf.movie_id = fr.movie_id
				AND pf.user_id = fr.user_id
				LEFT JOIN fr_attribute fa ON fa.fr_id = fr.id
				ORDER BY
				fa.like DESC,
				fr.id ASC
				LIMIT ?,? `
	o := orm.NewOrm()
	_, err = o.Raw(sql, pageIndex, pageSize).QueryRows(&fr)
	fmt.Println(fr[1].Id)
	return
}