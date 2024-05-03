package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `gorm:"name;not null"`
	Email          string `gorm:"column:email;unique;not null"`
	NoPhone        string `gorm:"column:no_phone;not null"`
	Password       string `gorm:"column:password;not null"`
	ProfilePicture string `gorm:"column:profile_picture;not null"`
}
