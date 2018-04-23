package models

import (
	"zcm_tools/orm"
	"fmt"
)

//历史记录
type HistoryRecord struct {
	Id      int
	Uid     string
	MovieId int
	Date    string
	MovieImg string `description:"电影图片"`
	MovieName string
}

//添加一条观看记录
func InsertHistoryRecord(h HistoryRecord)(err error){
	sql := `INSERT INTO history_record (uid,movie_id,date) VALUES(?,?,NOW())`
	_,err = orm.NewOrm().Raw(sql,h.Uid,h.MovieId).Exec()
	return
}
//删除一条历史记录
func DeleteHistory(uid string,movie_id int)(err error){
	sql := `DELETE FROM history_record WHERE uid = ? AND movie_id = ?`
	_,err = orm.NewOrm().Raw(sql,uid,movie_id).Exec()
	return
}

//查看是否又该条记录
func CheckHistoryReocrd(uid string ,movie_id int)(h *HistoryRecord,err error){
	sql := `SELECT * FROM history_record WHERE uid = ? AND movie_id = ?`
	err = orm.NewOrm().Raw(sql,uid,movie_id).QueryRow(&h)
	return
}

//更新最新的观影记录
func UpdateHistroyRecord(uid string,movie_id int)(err error){
	sql := `UPDATE history_record SET date = NOW() WHERE uid = ? AND movie_id = ?`
	_,err = orm.NewOrm().Raw(sql,uid,movie_id).Exec()
	return
}

//查看观影列表
func HistoryReocrd(uid string,pageIndex,pageSize int)(h []HistoryRecord,err error){
	sql := `SELECT h.*,m.movie_img,m.movie_name FROM history_record h
			LEFT JOIN movie m on m.id = h.movie_id
			WHERE uid = ?
			ORDER BY date DESC
			LIMIT ?,? `
	_,err = orm.NewOrm().Raw(sql,uid,pageIndex,pageSize).QueryRows(&h)
	fmt.Println(sql,uid,h)
	return
}

//统计该用户的观看电影的次数（相同电影重复看只算一次）
func HistoryRecordCount(uid string)(count int,err error){
	sql := `SELECT count(1) FROM history_record WHERE uid = ?`
	err = orm.NewOrm().Raw(sql,uid).QueryRow(&count)
	return
}
