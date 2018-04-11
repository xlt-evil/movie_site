package models

import (
	"time"
	"zcm_tools/orm"
)

//消息
//官方消息
type OfficalMsg struct {
	Id      int
	Title   string    `description:"标题"`
	Content string    `description:"内容"`
	Date    time.Time `description:"具体内容"`
	Href    string    `description:"跳转路径暂定"`
}

func FindOfficalMsg(pageIndex,pageSize int)(list []OfficalMsg,err error){
	sql := `SELECT title,content FROM official_msg
			LIMIT ?,?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,pageIndex,pageSize).QueryRows(&list)
	return
}

func CountOfficalMsg()(count int,err error){
	sql := `SELECT COUNT(1) FROM official_msg`
	o := orm.NewOrm()
	err = o.Raw(sql).QueryRow(&count)
	return
}