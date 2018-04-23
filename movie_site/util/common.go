package util

import (
	"math/rand"
	"time"
	"fmt"
)

//通用方法

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

//随机取出电影id
func RandomId(length int)(ids []int){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(length)
	for i:= 0;i<8;i++ {
		fmt.Println("213213")
		num := r.Intn(length)

		for ;DistinceNum(num,ids);{//不能有重复的id
		fmt.Println("+++++++++")
			num = r.Intn(length)
		}
		ids = append(ids,num)
	}
	fmt.Println(ids)
	return ids
}

//查看是否重复的数字
func DistinceNum(num int,id []int)(ok bool){
	for i,_ := range id{
		if id[i] == num {
			return !ok
		}
	}
	return ok
}
