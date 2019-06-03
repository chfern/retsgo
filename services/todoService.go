package services

import (
	"github.com/fernandochristyanto/retsgo/app/core"
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/fernandochristyanto/retsgo/repositories"
)

type TodoService struct {
	repository *repositories.TodoRepository
	logger     *core.Log
}

func (todoService TodoService) GetAll() *[]models.Todo {
	return todoService.repository.GetAll()
}

func (todoService TodoService) GetByID(id int) *models.Todo {
	return todoService.repository.GetByID(id)
}

func NewTodoService(logger *core.Log, repository *repositories.TodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
		logger:     logger,
	}
}
