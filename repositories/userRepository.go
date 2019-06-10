package repositories

import (
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/fernandochristyanto/retsgo/services"
)

type UserRepository struct {
	service *services.UserService
}

func (repo UserRepository) GetByUsernameAndPassword(username string, password string) *models.User {
	return repo.service.GetByUsernameAndPassword(username, password)
}

func NewUserRepository(service *services.UserService) *UserRepository {
	return &UserRepository{
		service: service,
	}
}
