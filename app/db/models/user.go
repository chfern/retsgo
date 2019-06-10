package models

type User struct {
	Base
	Username string `gorm:"type:varchar(255);column:Username" json:"username"`
	Password string `gorm:"type:varchar(255);column:Password" json:"-"`
	Todos    []Todo `gorm:"foreignkey:UserID;association_foreignkey:ID" json:"todos"`
}

func (User) TableName() string {
	return "users"
}
