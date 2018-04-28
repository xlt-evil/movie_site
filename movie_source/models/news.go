package models

import (
	"zcm_tools/orm"
	"fmt"
)

type News struct {
	Id      int
	Title   string `description:"标题"`
	Content string `description:"内容"`
	Img     string `description:"宣传海报"`
	Point   string `description:"简要"`
	Date    string
}

//官方消息
type OfficalMsg struct {
	Id      int
	Title   string    `description:"标题"`
	Content string    `description:"内容"`
	Date   	string    `description:"具体内容"`
	Href    string    `description:"跳转路径暂定"`
}

func InsertNews(n News)(err error){
	o := orm.NewOrm()
	defer func(){
		if err != nil {
			o.Rollback()
			fmt.Println(err.Error())
			return
		}
		o.Commit()
	}()
	sql := `INSERT INTO news (title,date,content,img,point)
			VALUES
			(?,NOW(),?,?,?)`
	_,err = o.Raw(sql,n.Title,n.Content,n.Img,n.Point).Exec()
	if err != nil {
		return
	}
	sql = `SELECT id FROM news ORDER BY id DESC LIMIT 1`
	id := 0
	err = o.Raw(sql).QueryRow(&id)
	if err != nil {
		return
	}
	err = InsertNewsAttribute(id,o)
	if err != nil {
		return
	}
	return
}

//新闻列表
func NewsList(pageIndex, pageSize int)(n []News,err error){
	sql := `SELECT id,title,date,content,point FROM news  ORDER BY date DESC LIMIT ?,? `
	_, err = orm.NewOrm().Raw(sql, pageIndex, pageSize).QueryRows(&n)
	return
}

//新闻总条数
func NewsListCount()(count int,err error){
	sql := `SELECT COUNT(1) FROM news `
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//得到图片
func GetNewsImg(id int)(img string,err error){
	sql := `SELECT img FROM news WHERE id = ? `
	err = orm.NewOrm().Raw(sql,id).QueryRow(&img)
	fmt.Println(err)
	return
}

//删除新闻
func DelNews(id int )(err error){
	sql := `DELETE FROM news WHERE id = ?`
	_,err = orm.NewOrm().Raw(sql,id).Exec()
	return
}

//新闻列表
func OfficalList(pageIndex,pageSize int)(o []OfficalMsg,err error){
	sql := `SELECT * FROM official_msg  ORDER BY date DESC LIMIT ?,?`
	_, err = orm.NewOrm().Raw(sql, pageIndex, pageSize).QueryRows(&o)

	return
}

func OfficalListCount()(count int,err error){
	sql := `SELECT COUNT(1) FROM official_msg  `
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//发布公告
func InsertOfficialMsg(title,content string)(err error){
	sql := `INSERT INTO official_msg (title,content,date,href) VALUES (?,?,NOW(),1)`
	_,err = orm.NewOrm().Raw(sql,title,content).Exec()
	fmt.Println(sql,title,content)
	return
}

func DelOfficialMsg(id int)(err error){
	sql := `DELETE FROM official_msg WHERE id = ?`
	_,err = orm.NewOrm().Raw(sql,id).Exec()
	return
}

//发布公告
func InsertOfficialMsgByMovie(title,content string,ormer orm.Ormer)(err error){
	sql := `INSERT INTO official_msg (title,content,date,href) VALUES (?,?,NOW(),1)`
	_,err = ormer.Raw(sql,title,content).Exec()
	fmt.Println(sql,title,content)
	return
}

func InsertNewsAttribute(newsid int,ormer orm.Ormer )(err error){
	sql := `INSERT INTO news_attribute (news_id) VALUES (?)`
	_,err = ormer.Raw(sql,newsid).Exec()
	return
}




