package services

import (
	"SB/movie_site/util"
	"SB/movie_site/models"
	"fmt"
)

//查询官方消息
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

func FindNewMsg(page int,pageSize int)(resp MovieResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,pageSize)
	n,err := models.GetNewsList(pageIndex,pageSize)
	if err != nil {
		resp.Msg = "查询新闻列表失败"
		return
	}
	for i,_ := range n{
		n[i].Img = util.News_img+n[i].Img
	}
	count,err := models.GetNewsCount()
	if err != nil {
		resp.Msg = "查询总条数失败"
		return
	}
	pages := util.PageNum(count,pageSize)
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
	resp.Object = n
	resp.Status = true
	resp.Page = page
	fmt.Println(resp.Object)
	return
}

func FindNewMsg2(page int,pageSize int)(resp MovieResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,pageSize)
	n,err := models.GetNewsListAsc(pageIndex,pageSize)
	if err != nil {
		resp.Msg = "查询新闻列表失败"
		return
	}
	for i,_ := range n{
		n[i].Img = util.News_img+n[i].Img
	}
	count,err := models.GetNewsCount()
	if err != nil {
		resp.Msg = "查询总条数失败"
		return
	}
	pages := util.PageNum(count,pageSize)
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
	resp.Object = n
	resp.Status = true
	resp.Page = page
	fmt.Println(resp.Object)
	return
}

func GetNewsDetail(id int )(resp MsgResponse,count int){
	n,err := models.GetNewsDetail(id)
	if err != nil{
		resp.Msg = "查看电影详情失败"
		return
	}
	n.Img = util.News_img+n.Img
	err = models.AddNewsVisiior(id)
	if err != nil {
		resp.Msg = "添加人气失败"
		return
	}
	count,err = models.GetVisitor(id)
	if err != nil {
		resp.Msg = "查看访问此数失败"
		return
	}
	resp.Object = n
	resp.Status = true

	return
}
