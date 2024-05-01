package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"name"`
	Email          string `gorm:"column:email"`
	NoPhone        string `gorm:"column:no_phone"`
	Password       string `gorm:"column:password"`
	ProfilePicture string `gorm:"column:profile_picture"`
}
