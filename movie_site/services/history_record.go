package services

import (
	"SB/movie_site/models"
	"SB/movie_site/util"
	"fmt"
)

//历史记录
func InsertHistory(uid string,movie_id int)(resp MsgResponse){
	resp.Status = true
	h,_ := models.CheckHistoryReocrd(uid,movie_id)
	if h == nil{
		err := models.InsertHistoryRecord(models.HistoryRecord{Uid:uid,MovieId:movie_id})
		if err != nil {
			resp.Msg = "添加用户观看记录失败"
			return
		}
	}else {
		err := models.UpdateHistroyRecord(uid,movie_id)
		if err != nil {
			resp.Msg = "更新用户最新的观看时间失败"
		}
	}
	resp.Status = true
	return
}

func GetUserHistoryRecord(uid string,page int)(resp MovieResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.UserCenterSize)
	h,_ := models.HistoryReocrd(uid,pageIndex,util.UserCenterSize)
	count,err :=  models.HistoryRecordCount(uid)
	if err != nil {
		resp.Msg = "统计观看记录失败"
		return
	}
	for i,_ := range h {
		h[i].MovieImg = util.Movie_img+h[i].MovieImg
	}
	pages := util.PageNum(count,util.UserCenterSize)
	ok := pagination(&resp,pages,count,page,util.UserCenterSize)
	if !ok {
		return
	}
	resp.Object = h
	resp.Status = true
	resp.Page = page
	if !resp.Next {
		resp.Page = 0
	}
	fmt.Println("xiuxiuxiux",resp)
	return
}

//删除我的历史记录
func DeleteMyHistory(uid string,movieId int)(resp MsgResponse){
	resp.Status = false
	err := models.DeleteHistory(uid,movieId)
	if err != nil {
		resp.Msg = "删除我的历史记录失败"
		return
	}
	resp.Status = true
	return
}