package repositories

import (
	"github.com/fernandochristyanto/retsgo/app/db/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (repo UserRepository) GetByUsernameAndPassword(username string, password string) *models.User {
	var user models.User
	repo.db.Where(&models.User{Username: username, Password: password}).First(&user)
	return &user
}

func (repo UserRepository) GetByID(id int64) *models.User {
	var user models.User
	repo.db.Find(&user, id)
	return &user
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
