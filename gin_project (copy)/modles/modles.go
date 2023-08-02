package modles

import "gin_project/database"

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	//
}

func CreateATodo(todo *ToDo) (err error) {
	err = database.DB.Create(&todo).Error
	return
}

func GetTodoList() (todoList []*ToDo, err error) {
	if err = database.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoById(id string) (todo *ToDo, err error) {
	todo = new(ToDo)
	if err = database.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *ToDo) (err error) {
	err = database.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = database.DB.Where("id=?", id).Delete(&ToDo{}).Error
	return
}
