package models

//连接数据库
import (
	"zcm_tools/orm"
	"SB/movie_site/util"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	orm.RegisterDataBase("default","mysql",util.MYSQL_URL)
}