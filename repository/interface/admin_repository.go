package _interface

import "mini_project/model/domain"

type AdminRepository interface {
	Register(admin *domain.Admin) (*domain.Admin, error)
	Login(username string) (*domain.Admin, error)
	GetByID(id int) (*domain.Admin, error)
}
