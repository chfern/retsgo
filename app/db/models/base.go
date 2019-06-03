package models

import (
	"time"
)

type Base struct {
	ID        int64      `gorm:"type:bigserial;primary_key;column:ID" json:"id"`
	CreatedAt time.Time  `gorm:"column:CreatedAt;" sql:"DEFAULT:NOW()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:UpdatedAt; null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:DeletedAt; null" sql:"index" json:"deleted_at"`
}
