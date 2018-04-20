package util


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
