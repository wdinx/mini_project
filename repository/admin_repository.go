package repository

import "mini_project/model/domain"

type AdminRepository interface {
	Register(admin *domain.Admin) (*domain.Admin, error)
	Login(username string) (*domain.Admin, error)
}
