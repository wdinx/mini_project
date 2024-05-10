package _interface

import "mini_project/model/domain"

type UserRepository interface {
	Login(email string) (*domain.User, error)
	Register(user *domain.User) error
}
