package service

import (
	"movie_source/models"
	"movie_source/util"
	"encoding/json"
	"time"
	"fmt"
	"strings"
)

//基本信息回复响应体
type MsgResp struct {
	Status bool `description:"业务状态"`
	Msg string  `description:"错误信息"`
	Object interface{} `description:"具体数据"`
}
//分页信息回复响应体
type PagemsgResp struct {
	MsgResp
	Pref bool //是否有上一页
	Next bool //是否有下一页
	Page int  //当前页
}

//登录
func Login(uid,password string)(resp MsgResp){
	resp.Status = false
	u,_ := models.CheckUser(uid,password)
	if u == nil {
		resp.Msg = "账户或密码错误"
		return
	}
	if (u.Status != 1){
		resp.Msg = "该账户被冻结！"
		return
	}
	b ,_ := json.Marshal(&u)
	//登录成功的用户添加到缓存区域
	if util.Re == nil{
		util.RedisPut(util.Rc,util.BaseKeyName+u.Uid,b,24*time.Hour)
	}
	resp.Status = true
	return
}
//登出
func LoginOut(key string)(resp MsgResp){
	if util.Re == nil{
		util.DelKey(util.Rc,util.BaseKeyName+key)
		fmt.Println("yes")
		resp.Status = true
		resp.Msg = "欢迎下次光临！"
	}else{
		resp.Status = false
		resp.Msg = "服务器异常，请稍后重试"
	}
	return
}

//今日数据查询
func CheckUserDate()(count []int){
	t ,err :=models.TodayRegisterCount()
	if err != nil {
		return
	}
	count = append(count,t)
	y,err := models.YesterdayRegisterCount()
	if err != nil {
		return
	}
	count = append(count,y)
	a,err := models.RegisterCount()
	if err != nil {
		return
	}
	count = append(count,a)
	return
}

//建议消息
func AdviseList(page int)(resp PagemsgResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.PageSize4)
	a,err := models.AdviseList(pageIndex,util.PageSize4)
	if err != nil {
		resp.Msg = "查询建议列表时出现bug了"
		return
	}
	count,err := models.AdviseListCount()
	if err != nil {
		resp.Msg = "查询总建议列表条数时出现bug了"
		return
	}
	pages := util.PageNum(count,util.PageSize4)
	pagination(&resp,pages,count,page,util.PageSize4)
	resp.Status = true
	resp.Page = page
	resp.Object = a
	fmt.Println(resp)
	return
}

//建议消息
func ResourceList(page int)(resp PagemsgResp){
	resp.Status = false
	pageIndex := util.PageIndex(page,util.PageSize4)
	a,err := models.ResourceList(pageIndex,util.PageSize4)
	if err != nil {
		resp.Msg = "查询建议列表时出现bug了"
		return
	}
	count,err := models.ResourceListCount()
	if err != nil {
		resp.Msg = "查询总建议列表条数时出现bug了"
		return
	}
	pages := util.PageNum(count,util.PageSize4)
	pagination(&resp,pages,count,page,util.PageSize4)
	resp.Status = true
	resp.Page = page
	resp.Object = a
	fmt.Println(resp)
	return
}

func MsgDetail(id int )(resp MsgResp){
	resp.Status = false

	m,err := models.MsgDetail(id)
	if err != nil {
		resp.Msg = "bug"
		return
	}
	if m.Status != "pass" {
		s,err := models.GetTemplate(1)
		if err != nil{
			resp.Msg = "获取模板失败"
			return
		}
		content := strings.Replace(s.Content,"content",m.Title,1)
		fmt.Println(s.Content,content)
		err = models.InsertMsgRecord(models.MsgRecord{
			UserId:m.Uid,
			Title:s.Title,
			Content:content,
		})
		if err != nil {
			resp.Msg = "发送消息失败"
			return
		}
	}
	err = models.UpdateState(id)
	if err != nil {
		resp.Msg = "查看消息状态失败"
		return
	}
	resp.Status = true
	resp.Object = m
	return
}

//分页方法
func pagination(resp *PagemsgResp,pages,count,page,pageSize int)(bool){
	if page < pages {
		resp.Next = true
	}
	if count >pageSize && page != 1{
		resp.Pref = true
	}
	if count == 0 {
		resp.Msg = "没有数据"
		b := false
		return b
	}
	return true
}


