package service

import (
	"movie_source/models"
	"movie_source/util"
)

//添加用户
func AddAdmin(role int,uid,password string)(resp MsgResp){
	if role != 1 {
		resp.Msg = "你没有权限添加管理员"
		return
	}
	u,_ := models.IsUser(uid)
	if u != nil {
		resp.Msg = "已存在该用户"
		return
	}
	err := models.InsertAdmin(uid,password)
	if err != nil {
		resp.Msg = "添加系统用户失败"
		return
	}
	resp.Status = true
	return
}

//用户列表
func AdminList(page int)(resp PagemsgResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.PageSize4)
	a ,err := models.AdminList(pageIndex,util.PageSize4)
	if err != nil {
		resp.Msg = "查询建议列表时出现bug了"
		return
	}
	count,err := models.AdminListCount()
	if err != nil {
		resp.Msg = "查询总建议列表条数时出现bug了"
		return
	}
	pages := util.PageNum(count,util.PageSize4)
	pagination(&resp,pages,count,page,util.PageSize4)
	resp.Status = true
	resp.Page = page
	resp.Object = a
	return
}

//冻结账户
func Frozen(role,id,status int )(resp MsgResp){
	resp.Status = false
	if role != 1 {
		resp.Msg = "你没有该权限"
		return
	}
	err := models.ForzenAdimin(id,status)
	if err != nil {
		resp.Msg = "冻结失败"
		return
	}
	resp.Status = true
	resp.Object = status
	return
}

//删除管理员
func DelAdmin(role,id int)(resp MsgResp){
	resp.Status = false
	if role != 1 {
		resp.Msg = "你没有该权限"
		return
	}
	err := models.DelAdmin(id)
	if err != nil {
		resp.Msg = "删除用户失败"
		return
	}
	resp.Status = true
	return
}


