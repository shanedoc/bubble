package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r:=gin.Default()
	//引入css、js文件
	r.Static("/static","static")
	//模版文件路径
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandle)

	//v1
	v1Group:=r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看:所有or一个
		v1Group.GET("/todo", controller.GetTodoList)
		/*v1Group.GET("/todo/:id", func(c *gin.Context) {
			id,ok:=c.Params.Get("id")
			if !ok{
				c.JSON(http.StatusOK,gin.H{"error":"id有误"})
				return
			}
			var todo Todo
			if err=DB.Where("id=?",id).First(&todo).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
				return
			}else{
				c.JSON(http.StatusOK,todo)
			}
		})*/
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
	}
	return r
}