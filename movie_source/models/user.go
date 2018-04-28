package models

import "zcm_tools/orm"



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

//模板
type Template struct {
	Title   string
	Content string
}

type MsgRecord struct {
	UserId  string
	Title   string
	Content string
}



//今日用户注册
func TodayRegisterCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM user WHERE register_time = CURDATE()`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//昨日用户注册
func YesterdayRegisterCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM user WHERE register_time = CURDATE()-1`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//总人数
func RegisterCount() (count int, err error) {
	sql := `SELECT COUNT(1) FROM user `
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//建议消息
func AdviseList(pageIndex, pageSize int) (a []Advise, err error) {
	sql := `SELECT * FROM advise WHERE role = 'build' AND status ='wait' ORDER BY advise_date DESC , id LIMIT ?,? `
	_, err = orm.NewOrm().Raw(sql, pageIndex, pageSize).QueryRows(&a)
	return
}

//建议消息总条数
func AdviseListCount() (count int, err error) {
	sql := `SELECT count(1) FROM advise WHERE role = 'build' AND status ='wait' `
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//建议消息
func ResourceList(pageIndex, pageSize int) (a []Advise, err error) {
	sql := `SELECT * FROM advise WHERE role = 'resource' AND status ='wait' ORDER BY advise_date DESC , id LIMIT ?,? `
	_, err = orm.NewOrm().Raw(sql, pageIndex, pageSize).QueryRows(&a)
	return
}

//建议消息总条数
func ResourceListCount() (count int, err error) {
	sql := `SELECT count(1) FROM advise WHERE role = 'resource' AND status ='wait' `
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//消息具体内容
func MsgDetail(id int) (a Advise, err error) {
	sql := `SELECT * FROM advise WHERE id = ? `
	err = orm.NewOrm().Raw(sql, id).QueryRow(&a)
	return
}

//更新消息查看状态
func UpdateState(id int) (err error) {
	sql := `UPDATE advise SET status = 'pass' WHERE id = ?`
	_, err = orm.NewOrm().Raw(sql, id).Exec()
	return
}

//获取模板
func GetTemplate(id int) (t Template, err error) {
	sql := `SELECT content,title FROM template WHERE id = ?`
	err = orm.NewOrm().Raw(sql, id).QueryRow(&t)
	return
}


//添加用户消息
func InsertMsgRecord(m MsgRecord) (error){
	sql := `INSERT INTO msg_record 	(user_id,title,content,status) VALUES (?,?,?,'false')`
	_,err := orm.NewOrm().Raw(sql,m.UserId,m.Title,m.Content).Exec()
	return err
}


