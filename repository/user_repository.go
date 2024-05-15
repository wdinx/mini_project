package repository

import "mini_project/model/domain"

type UserRepository interface {
	Login(email string) (*domain.User, error)
	Register(user *domain.User) error
	GetByID(userID int) (*domain.User, error)
}
