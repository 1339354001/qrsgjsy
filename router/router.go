package router

import (
	"gin_project/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")

	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.Update)
		// 删除
		v1Group.DELETE("/todo/:id", controller.Delete)
	}
	return r
}
