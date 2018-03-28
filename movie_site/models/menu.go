package models

import (
	"zcm_tools/orm"
)

//基本的菜单栏渲染

type Menu struct {
	Id int `description:"唯一编号"`
	Name string `description:"菜单名"`
}



func FindAreaMenu()(a []Menu,err error){
	sql := `SELECT * From area`
	_,err = orm.NewOrm().Raw(sql).QueryRows(&a)
	return
}


func FindTypeMenu()(a []Menu,err error){
	sql := `SELECT * From type`
	_,err = orm.NewOrm().Raw(sql).QueryRows(&a)
	return
}
