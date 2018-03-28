package models

import (
	"zcm_tools/orm"
	"fmt"
	"strconv"
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

//添加账户
func AddUser(u User)error{
	sql := `INSERT INTO user (uid,password,question,answer) VALUES(?,?,?,?)`
	o := orm.NewOrm()
	_,err := o.Raw(sql,u.Uid,u.Password,u.Question,u.Answer).Exec()
	return err
}

//查用户id是否存在
func CheckUid(uid string)(name string){
	sql := `SELECT uid FROM user WHERE uid = ?`
	o := orm.NewOrm()
	o.Raw(sql,uid).QueryRow(&name)
	return
}

//判断用户账户密码

func CheckPwdAndUid(u User)(user *User){
	sql := `SELECT * FROM user WHERE uid = ? AND password = ? `
	o := orm.NewOrm()
	o.Raw(sql,u.Uid,u.Password).QueryRow(&user)
	fmt.Print(user)
	return
}

func ForgetPwd(u User)(password string){
	sql := `SELECT password FROM user WHERE uid = ? AND question = ? AND answer = ?`
	o := orm.NewOrm()
	question ,_:= strconv.Atoi(u.Question)
	o.Raw(sql,u.Uid,question,u.Answer).QueryRow(&password)
	return
}