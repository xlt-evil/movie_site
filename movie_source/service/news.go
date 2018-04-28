package service

import (
	"movie_source/models"
	"movie_source/util"
)

//发布新闻
func ReleaseNews(titie,point,content,name string)(resp MsgResp){
	resp.Status = false
	err := models.InsertNews(models.News{
		Title:titie,
		Content:content,
		Point:point,
		Img:name,
	})
	if err != nil {
		resp.Msg = "发布新闻失败"
		return
	}
	resp.Status = true
	return
}
//新闻列表
func NewsList(page int )(resp PagemsgResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.PageSize4)
	n,err := models.NewsList(pageIndex,util.PageSize4)
	if err != nil {
		resp.Msg = "查询建议列表时出现bug了"
		return
	}
	count,err := models.NewsListCount()
	if err != nil {
		resp.Msg = "查询总建议列表条数时出现bug了"
		return
	}
	pages := util.PageNum(count,util.PageSize4)
	pagination(&resp,pages,count,page,util.PageSize4)
	resp.Status = true
	resp.Page = page
	resp.Object = n
	return
}
//删除新闻
func DelNews(id int)(resp MsgResp){
	resp.Status = false
	img ,err := models.GetNewsImg(id)
	if err != nil {
		resp.Msg = "获取图片失败"
		return
	}
	err = util.DelImg(img)
	if err != nil {
		resp.Msg = err.Error()
	}
	err = models.DelNews(id)
	if err != nil {
		resp.Msg = err.Error()
	}
	resp.Status = true
	return
}
//获取官方消息
func OfficalList(page int)(resp PagemsgResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.PageSize4)
	o,err := models.OfficalList(pageIndex,util.PageSize4)
	if err != nil {
		resp.Msg = "查询建议列表时出现bug了"
		return
	}
	count,err := models.OfficalListCount()
	if err != nil {
		resp.Msg = "查询总建议列表条数时出现bug了"
		return
	}
	pages := util.PageNum(count,util.PageSize4)
	pagination(&resp,pages,count,page,util.PageSize4)
	resp.Status = true
	resp.Page = page
	resp.Object = o
	return
}
//发布官方消息
func ReleaseOfficialMsg(title,content string)(resp MsgResp){
	resp.Status = true
	err := models.InsertOfficialMsg(title,content)
	if err != nil {
		resp.Msg = "发布新闻失败"
		return
	}
	resp.Status = true
	return
}
//删除信息
func DelOfficialMsg(id int )(resp MsgResp){
	resp.Status = true
	err := models.DelOfficialMsg(id)
	if err != nil {
		resp.Msg = "删除电影失败"
		return
	}
	resp.Status = true
	return
}



