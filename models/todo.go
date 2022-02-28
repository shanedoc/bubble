package models

import (
	"bubble/dao"
	"fmt"
)

//todo model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

//todo model curd
//一般公司的业务逻辑：controller->service(logic)->model
func CreateATodo(todo *Todo)(err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList() (todoList []*Todo,err error) {
	if err=dao.DB.Find(&todoList).Error;err!=nil{
		return nil,err
	}
	fmt.Println(todoList)
	return
}

func GetATodo(id string)(todo *Todo,err error)  {
	if err=dao.DB.Where("id=?",id).First(todo).Error;err!=nil{
		return nil,err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err=dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string)(err error)  {
	err=dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}