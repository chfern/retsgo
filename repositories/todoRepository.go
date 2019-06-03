package repositories

import (
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func (repo TodoRepository) GetAll() *[]models.Todo {
	var todos []models.Todo
	repo.db.Find(&todos)
	return &todos
}

func (repo TodoRepository) GetByID(id int) *models.Todo {
	var todo models.Todo
	repo.db.Find(&todo, id)
	return &todo
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}
