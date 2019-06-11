package services

import (
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (service UserService) GetByUsernameAndPassword(username string, password string) *models.User {
	var user models.User
	service.db.Where(&models.User{Username: username, Password: password}).First(&user)
	return &user
}

func (service UserService) GetByID(id int64) *models.User {
	var user models.User
	service.db.Find(&user, id)
	return &user
}

func (service UserService) GetTodos(id int64) *[]models.Todo {
	user := service.GetByID(id)
	var todos []models.Todo
	a := service.db.Model(&user)
	a.Related(&todos)
	return &todos
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}