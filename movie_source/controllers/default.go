package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"strings"
	"io"
	"fmt"
)

type MainController struct {
	beego.Controller
}


//网站传过来的图片
func (this *MainController)UpdateImg(){
	flag := "0"
	newname := ""
	uid := this.GetString("uid")
	defer func() {
		this.Redirect("http://localhost:9090/user/updateimg?flag="+flag+"&img="+newname+"&uid="+uid,302)
	}()
	s,f,err:= this.GetFile("file")
	if err != nil {
		return
	}
	name := f.Filename
	mode := strings.Split(name,".")
	newname = uid+"."+mode[len(mode)-1]
	os.Remove("source/Img/per_img/"+newname)
	f1 ,err:= os.Create("source/Img/per_img/"+newname)
	fmt.Println(f1.Name())
	if err != nil {
		return
	}
	io.Copy(f1,s)
	s.Close()
	f1.Close()
	flag = "1"
	fmt.Println("ok")
}







