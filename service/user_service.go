package service

import (
	"mini_project/model/web"
)

type UserService interface {
	Login(request *web.UserLoginRequest) (*web.UserLoginResponse, error)
	Register(request *web.UserRegisterRequest) (*web.UserResponse, error)
}
