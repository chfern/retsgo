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

func (repo UserRepository) GetByID(id int64) *models.User {
	return repo.service.GetByID(id)
}

func NewUserRepository(service *services.UserService) *UserRepository {
	return &UserRepository{
		service: service,
	}
}
