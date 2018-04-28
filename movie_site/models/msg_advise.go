package models

import (
	"fmt"
	"time"
	"zcm_tools/orm"
)

//消息 建议
//官方消息
type OfficalMsg struct {
	Id      int
	Title   string    `description:"标题"`
	Content string    `description:"内容"`
	Date    time.Time `description:"具体内容"`
	Href    string    `description:"跳转路径暂定"`
}

//建议表
type Advise struct {
	Id         int
	Uid        string
	Role       string `description:"建议类型"`
	Title      string `description:"标题"`
	Content    string `description:"内容"`
	Status     string `description:"查看状态"`
	AdviseDate string `description:"建议时间"`
}

//新闻表
type News struct {
	Id      int
	Title   string `description:"标题"`
	Content string `description:"内容"`
	Img     string `description:"宣传海报"`
	Point   string `description:"简要"`
	Date    string
}

//新闻属性表
type NewsAttribute struct {
	Id      int
	NewsId  int
	Visitor int `description:"访问人数"`
}

//系统通知
type SystemMsg struct {
	Id      int
	UserId  int
	Title   string
	Content string
	Status  string
}

//找到官方消息
func FindOfficalMsg(pageIndex, pageSize int) (list []OfficalMsg, err error) {
	sql := `SELECT title,content FROM official_msg
			LIMIT ?,?`
	o := orm.NewOrm()
	_, err = o.Raw(sql, pageIndex, pageSize).QueryRows(&list)
	return
}

//统计官方消息
func CountOfficalMsg() (count int, err error) {
	sql := `SELECT COUNT(1) FROM official_msg`
	o := orm.NewOrm()
	err = o.Raw(sql).QueryRow(&count)
	return
}

//插入一条新的建议
func InsertAdvise(a Advise) (err error) {
	sql := `INSERT advise (uid,role,title,content,status,advise_date) VALUES (?,?,?,?,'wait',NOW())`
	_, err = orm.NewOrm().Raw(sql, a.Uid, a.Role, a.Title, a.Content).Exec()
	return
}

//找到建议列表
func FindAdviseList(uid string, pageIndex, pageSize int) (a []Advise, err error) {
	sql := `SELECT * FROM advise
			WHERE uid = ?
			ORDER BY advise_date DESC
			LIMIT ?,?`
	_, err = orm.NewOrm().Raw(sql, uid, pageIndex, pageSize).QueryRows(&a)
	return
}

//找到建议列表总数
func FindAdviseCount(uid string) (count int, err error) {
	sql := `SELECT COUNT(1) From advise
			WHERE uid = ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&count)
	return
}

//首页官网查询
func GetIndexOffical() (list []OfficalMsg, err error) {
	sql := `SELECT title FROM official_msg
			ORDER BY date DESC
			LIMIT 0,10`
	_, err = orm.NewOrm().Raw(sql).QueryRows(&list)
	return
}

//新闻列表
func GetNewsList(page, pageSize int) (n []News, err error) {
	sql := `SELECT * FROM news ORDER BY id DESC LIMIT ?,?`
	_, err = orm.NewOrm().Raw(sql, page, pageSize).QueryRows(&n)
	fmt.Println(n)
	return
}

//新闻总数
func GetNewsCount() (count int, err error) {
	sql := `SELECT count(1) FROM news `
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//新闻列表
func GetNewsListAsc(page, pageSize int) (n []News, err error) {
	sql := `SELECT n.* FROM news n
			LEFT JOIN news_attribute na ON n.id = na.news_id
			ORDER BY na.visitor DESC
			LIMIT ?,?`
	_, err = orm.NewOrm().Raw(sql, page, pageSize).QueryRows(&n)
	fmt.Println(n)
	return
}

func GetNewsDetail(id int) (n News, err error) {
	sql := `SELECT * FROM news WHERE id = ?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&n)
	fmt.Println(n)
	return
}

func AddNewsVisiior(id int) (err error) {
	sql := `UPDATE news_attribute SET visitor = visitor+1 WHERE news_id = ?`
	_, err = orm.NewOrm().Raw(sql, id).Exec()
	return
}

func GetVisitor(id int) (visitor int, err error) {
	sql := `SELECT visitor FROM news_attribute  WHERE news_id = ?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&visitor)
	return
}

func SystemList(uid string ,pageIndex, pageSize int )(list []SystemMsg,err error){
	sql := `SELECT * FROM msg_record WHERE user_id = ? ORDER BY id DESC LIMIT ?,? `
	_,err = orm.NewOrm().Raw(sql,uid,pageIndex,pageSize).QueryRows(&list)
	fmt.Println(sql,uid,pageIndex,pageSize)
	return
}

func SystemListCount(uid string)(count int ,err error){
	sql := `SELECT COUNT(1) FROM msg_record WHERE user_id = ? `
	err = orm.NewOrm().Raw(sql,uid).QueryRow(&count)
	fmt.Println("jadsojasidjsaudnjus:        ",count)
	return
}
