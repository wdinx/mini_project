package service

import "mini_project/model/web"

type AdminService interface {
	Register(admin *web.AdminRegisterRequest) (*web.AdminRegisterResponse, error)
	Login(admin web.AdminLoginRequest) (*web.AdminLoginResponse, error)
}
