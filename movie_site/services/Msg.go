package services

import (
	"SB/movie_site/util"
	"SB/movie_site/models"
)

func FindOfficalMsg(page int )(resp MovieResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.OfficalSize)
	m,err := models.FindOfficalMsg(pageIndex,util.OfficalSize)
	if err != nil {
		resp.Msg = "查询官方消息失败"
		return
	}
	count,err := models.CountOfficalMsg()
	if err != nil {
		resp.Msg = "查询总条数失败"
		return
	}
	pages := util.PageNum(count,util.OfficalSize)
	if page <pages {
		resp.Next = true
	}
	if count > 3 && page != 1 {
		resp.Pref = true
	}
	if count == 0 {
		resp.Msg = "暂时没有官方消息"
		return
	}
	resp.Object = m
	resp.Status = true
	resp.Page = page
	return
}
