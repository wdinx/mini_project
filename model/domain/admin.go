package domain

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
