package controller

import (
	"gin_project/modles"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 挥发请求到这里
	// 从请求中把数据拿出来
	var todo modles.ToDo
	c.BindJSON(&todo)
	// 存入数据库
	err := modles.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	// 返回响应
}

func GetTodoList(c *gin.Context) {
	// 存入数据库
	todoList, err := modles.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
	// 返回响应
}

func Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	todo, err := modles.GetTodoById(id)
	if todo.Status == false {
		now := time.Now()
		todo.Time = now.Format("签到时间15:04:05")
	} else {
		todo.Time = ""
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
	c.BindJSON(&todo)
	err = modles.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func Delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	err := modles.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": "delete"})
	}
}
