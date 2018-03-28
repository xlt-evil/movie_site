package util

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"

)

//各种文件解析的方法

//redis 的解析
func RedisParse(config string)(re redis.Conn, err error){
	var c map[string]string

	json.Unmarshal([]byte(config),&c)
	if _, ok := c["key"];!ok{
		panic("redis配置文件出错了！")
	}
	re,err= redis.Dial("tcp",c["key"])
	if err != nil {
		panic("redis初始化失败")
	}
	_,err = re.Do("auth",c["password"])
	return
}
