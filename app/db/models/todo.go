package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	// Base
	UserID    uint   `gorm:"column:UserID" json:"user_id"`
	TaskName  string `gorm:"type:varchar(255);column:TaskName" json:"task_name"`
	Completed bool   `gorm:"type:bool;column:Completed" json:"completed"`
}

func (Todo) TableName() string {
	return "todos"
}
