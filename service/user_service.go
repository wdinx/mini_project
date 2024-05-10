package service

import (
	"mini_project/model/domain"
	"mini_project/model/web"
)

type UserService interface {
	Login(request *web.UserLoginRequest) (*web.UserResponse, error)
	Register(request *web.UserRegisterRequest) (*domain.User, error)
}
