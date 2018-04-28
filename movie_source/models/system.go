package models

import "zcm_tools/orm"

//管理员表
type Admin struct {
	Id       int
	Uid      string
	Password string
	Role     int
	Status   int
}

//用户密码检查
func CheckUser(uid, password string) (u *Admin, err error) {
	sql := `SELECT * FROM admin WHERE uid = ? AND password = ?`
	err = orm.NewOrm().Raw(sql, uid, password).QueryRow(&u)
	return
}

//添加用户
func InsertAdmin(uid,password string)(err error){
	sql := `INSERT INTO admin (uid,password,role,status) VALUES (?,?,2,1)`
	_,err = orm.NewOrm().Raw(sql,uid,password).Exec()
	return
}

//冻结用户
func ForzenAdimin(uid int,status int)(err error){
	sql := `UPDATE admin SET status = ? WHERE id = ? `
	_,err = orm.NewOrm().Raw(sql,status,uid).Exec()
	return
}

//管理员列表
func AdminList(pageIndex,pageSize int)(a []Admin,err error){
	sql := `SELECT * FROM admin WHERE role=2 LIMIT ?,? `
	_,err = orm.NewOrm().Raw(sql,pageIndex,pageSize).QueryRows(&a)
	return
}

//管理员列表总条数
func AdminListCount()(count int ,err error){
	sql := `SELECT COUNT(1) FROM admin WHERE role = 2 `
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

//判读是否有用户
func IsUser(uid string) (u *Admin, err error) {
	sql := `SELECT * FROM admin WHERE uid = ? `
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&u)
	return
}

func DelAdmin(id int)(err error){
	sql := `DELETE FROM admin WHERE id = ?`
	_,err = orm.NewOrm().Raw(sql,id).Exec()
	return
}