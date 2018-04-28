package main

import (
	_ "movie_source/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("perimg","source/Img/per_img")
	beego.SetStaticPath("movieimg","source/Img/movie_img")
	beego.SetStaticPath("vedio","source/video")
	beego.SetStaticPath("news","source/Img/news_img")
	beego.Run()
}

