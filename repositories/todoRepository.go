package repositories

import (
	"github.com/fernandochristyanto/retsgo/app/core"
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/fernandochristyanto/retsgo/services"
)

type TodoRepository struct {
	service *services.TodoService
	logger  *core.Log
}

func (repo TodoRepository) GetAll() *[]models.Todo {
	return repo.service.GetAll()
}

func (repo TodoRepository) GetByID(id int) *models.Todo {
	return repo.service.GetByID(id)
}

func NewTodoRepository(service *services.TodoService, logger *core.Log) *TodoRepository {
	return &TodoRepository{
		logger:  logger,
		service: service,
	}
}
