package util

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

//redis的一些常用命令的封装

func IsExist(rc redis.Conn,key string)bool{
	v,err := redis.Bool(rc.Do("EXISTS",key))
	if err != nil{
		return false
	}
	return v
}

//可以设置过期时间，覆盖旧值
func RedisPut (rc redis.Conn,key string,value []byte,timeout time.Duration)error{
	ok,err := rc.Do("SETEX",key,int64(timeout/time.Second),value)
	fmt.Print(ok)
	return err
}

//删除可以key
func DelKey(rc redis.Conn ,key string)error{
	var err error
	if _, err = rc.Do("DEL", key); err != nil {
		return err
	}
	return err
}

//读取redis 的数据
func RedisRead(rc redis.Conn ,key string)(data []byte,err error){
	data,err = redis.Bytes(Get(rc,key),err)
	return
}

// Get cache from redis.
func Get(rc redis.Conn,key string) interface{} {
	if v, err := rc.Do("GET", key); err == nil {
		return v
	}
	return nil
}
