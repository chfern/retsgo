package repositories

import (
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/fernandochristyanto/retsgo/services"
)

type TodoRepository struct {
	service *services.TodoService
}

func (repo TodoRepository) GetAll() *[]models.Todo {
	return repo.service.GetAll()
}

func (repo TodoRepository) GetByID(id int) *models.Todo {
	return repo.service.GetByID(id)
}

func NewTodoRepository(service *services.TodoService) *TodoRepository {
	return &TodoRepository{
		service: service,
	}
}
