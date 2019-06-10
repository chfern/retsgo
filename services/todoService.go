package services

import (
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/jinzhu/gorm"
)

type TodoService struct {
	db *gorm.DB
}

func (service TodoService) GetAll() *[]models.Todo {
	var todos []models.Todo
	service.db.Find(&todos)
	return &todos
}

func (service TodoService) GetByID(id int) *models.Todo {
	var todo models.Todo
	service.db.Find(&todo, id)
	return &todo
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}
