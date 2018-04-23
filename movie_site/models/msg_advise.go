package models

import (
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
	_,err = orm.NewOrm().Raw(sql,a.Uid,a.Role,a.Title,a.Content).Exec()
	return
}

//找到建议列表
func FindAdviseList(uid string, pageIndex ,pageSize int)(a []Advise,err error){
	sql := `SELECT * FROM advise
			WHERE uid = ?
			ORDER BY advise_date DESC
			LIMIT ?,?`
	_,err = orm.NewOrm().Raw(sql,uid,pageIndex,pageSize).QueryRows(&a)
	return
}

//找到建议列表总数
func FindAdviseCount(uid string)(count int ,err error){
	sql := `SELECT COUNT(1) From advise
			WHERE uid = ?`
	err = orm.NewOrm().Raw(sql,uid).QueryRow(&count)
	return
}

