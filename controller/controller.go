package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}

func CreateTodo(c *gin.Context) {
	//获取提交数据入库
	var todo models.Todo
	c.BindJSON(&todo)
	err:=models.CreateATodo(&todo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}

func GetTodoList(c *gin.Context) {
	todoList,err:=models.GetTodoList()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}else{
		c.JSON(http.StatusOK,todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{"error":"id有误"})
		return
	}
	todo,err:=models.GetATodo(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err=models.UpdateATodo(todo);err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}else{
		c.JSON(http.StatusOK,todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id,ok:=c.Params.Get("id")
	if !ok{
		c.JSON(http.StatusOK,gin.H{"error":"id有误"})
		return
	}
	err:=models.DeleteATodo(id)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}else{
		//c.JSON(http.StatusOK,todo)
		c.JSON(http.StatusOK,gin.H{"id":"删除成功"})
	}
}