package util

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

//配置表信息的初始化

var (
	RunMode     string     //运行模式
	MYSQL_URL   string     //主库
	Rc          redis.Conn //Redis
	Re          error      //Redis 错误
	RedisConfig string     //redis 配置
	Person_img  string     //头像的存储前缀
	Movie_img string//电影海报前缀
	Source_path string //电影资源
)

func init() {
	//读取配置表信息

	RunMode = beego.AppConfig.String("run_mode")
	//得到runmode下的配置信息
	config, err := beego.AppConfig.GetSection(RunMode)
	if err != nil {
		panic("数据库错误" + err.Error())
	}
	MYSQL_URL = config["mysql_url"]
	RedisConfig = config["redis_config"]
	Person_img = config["person_img"]
	Movie_img = config["movie_img"]
	Source_path = config["movie_vedio"]
	Rc, Re = RedisParse(RedisConfig)
	beego.Info("----------------")
	beego.Info("运行模式", RunMode)
	beego.Info(Source_path,"13fewhwefuh")
	beego.Info("----------------")
}
