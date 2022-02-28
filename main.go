package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main()  {
	//创建数据库、数据表
	//连接数据路
	err:=dao.InitMysql()
	if err!=nil{
		fmt.Println(err)
		panic(err)
		//todo log
	}
	//关闭
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	//注册路由
	fmt.Println("main")
	r:=routers.SetUpRouter()
	//不写端口默认8080
	r.Run()
}
