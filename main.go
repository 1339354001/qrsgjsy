package main

import (
	"gin_project/database"
	"gin_project/router"
)

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	Time   string `json:"time"`
	//
}

func main() {
	err := database.InitMySQL()
	if err != nil {
		panic(err)
	}
	err = database.DB.AutoMigrate(&ToDo{})
	if err != nil {
		panic(err)
	}
	r := router.SetUpRouter()
	r.Run(":9090")
}
