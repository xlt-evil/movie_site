package models

import (
	"zcm_tools/orm"
	"SB/movie_site/util"
	_ "github.com/go-sql-driver/mysql"
)

//连接数据库

func init(){
	orm.RegisterDataBase("default","mysql",util.MYSQL_URL)
}

