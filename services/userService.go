package services

import (
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/fernandochristyanto/retsgo/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func (service UserService) GetByUsernameAndPassword(username string, password string) *models.User {
	return service.userRepository.GetByUsernameAndPassword(username, password)
}

func (service UserService) GetByID(id int64) *models.User {
	return service.userRepository.GetByID(id)
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
