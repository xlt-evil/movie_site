package util

import (
	"mime/multipart"
	"strings"
	"strconv"
	"os"
	"io"
	"fmt"
	"github.com/pkg/errors"
)

var (
	BaseKeyName   = "Movie_BackGROUNDSYSTEM" //redis的基本key前缀
	PageSize4 = 4
)

//条数起始下标
func PageIndex(page,pageSize int)(int){
	if page == 1 {
		page = 0
	}else {
		page = (page-1)*pageSize
	}
	return page
}

//总页数
func PageNum(count,pageSize int,)(a int){//总页数
	if count%pageSize != 0{
		return (count/pageSize)+1
	}
	return count/pageSize
}

////页面数
//func Pages(page int ,pages int)(p []int){
//	if pages < 4 {
//		for i:=1 ;i <= pages ;i++ {
//			p = append(p,i)
//		}
//		return
//	}
//	stop := page - 2
//	for i := page;i > stop ;i-- {
//		p = append(p,i)
//	}
//	fmt.Println("pages",pages)
//	return
//}

func SaveImg(s multipart.File,f *multipart.FileHeader,id int,path string)(name string){
	mode := strings.Split(f.Filename,".")
	newsname := strconv.Itoa(id)+"."+mode[len(mode)-1]
	fmt.Println(mode)
	fmt.Println(newsname)
	os.Remove("source/Img/"+path+newsname)
	f1,err := os.Create("source/Img/"+path+newsname)
	if err != nil {
		return
	}
	_,err = io.Copy(f1,s)
	if err != nil {
		fmt.Println(err.Error())
	}
	f1.Close()
	name = newsname
	return name
}

func DelImg(img string)(error){
	err := os.Remove("source/Img/news_img/"+img)
	if err != nil {
		return errors.New("删除文件失败")
	}
	return nil
}

func SaveVideo(s multipart.File,f *multipart.FileHeader,id int)(name string){
	mode := strings.Split(f.Filename,".")
	newsname := strconv.Itoa(id)+"."+mode[len(mode)-1]
	fmt.Println(mode)
	fmt.Println(newsname)
	os.Remove("source/Video/"+newsname)
	f1,err := os.Create("source/Video/"+newsname)
	if err != nil {
		return
	}
	_,err = io.Copy(f1,s)
	if err != nil {
		fmt.Println(err.Error())
	}
	f1.Close()
	name = newsname
	return name
}