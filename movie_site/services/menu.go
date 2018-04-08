package services

import "SB/movie_site/models"

func FindAreaMenu()(resp MsgResponse){
	a , err := models.FindAreaMenu()
	if err != nil {
		resp.Status = false
		resp.Msg = "菜单获取失败"
		return
	}
	resp.Status = true
	resp.Object = a
	return
}

func FindTypeMenu()(resp MsgResponse){
	a , err := models.FindTypeMenu()
	if err != nil {
		resp.Status = false
		resp.Msg = "菜单获取失败"
		return
	}
	resp.Status = true
	resp.Object = a
	return
}
