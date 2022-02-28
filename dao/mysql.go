package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//连接mysql

var(
	//全局变量
	DB *gorm.DB
)

func InitMysql()(err error) {
	dsn:="root:Wozuihaokan@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err =gorm.Open("mysql",dsn)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//测试db连接
	return DB.DB().Ping()
}

func Close()  {
	DB.Close()
}
