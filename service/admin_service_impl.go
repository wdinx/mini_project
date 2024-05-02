package service

import (
	"mini_project/constant"
	"mini_project/middleware"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util"
)

type AdminServiceImpl struct {
	adminRepository repository.AdminRepository
}

func NewAdminRepository(adminRepository repository.AdminRepository) AdminService {
	return &AdminServiceImpl{
		adminRepository: adminRepository,
	}
}

func (service AdminServiceImpl) Register(admin *web.AdminRegisterRequest) (*web.AdminRegisterResponse, error) {
	if admin.Name == "" || admin.Username == "" || admin.Password == "" {
		return &web.AdminRegisterResponse{}, constant.ErrEmptyInput
	}

	var err error
	admin.Password, err = util.HashPassword(admin.Password)
	util.PanicIfError(err)

	insertAdmin := domain.Admin{
		Name:     admin.Name,
		Username: admin.Username,
		Password: admin.Password,
	}

	adminResponse, err := service.adminRepository.Register(&insertAdmin)
	if err != nil {
		return &web.AdminRegisterResponse{}, constant.ErrInsertDatabase
	}

	result := web.AdminRegisterResponse{
		Name:     adminResponse.Name,
		Username: adminResponse.Username,
		Password: adminResponse.Password,
	}

	return &result, nil
}

func (service AdminServiceImpl) Login(admin web.AdminLoginRequest) (*web.AdminLoginResponse, error) {
	if admin.Username == "" || admin.Password == "" {
		return &web.AdminLoginResponse{}, constant.ErrEmptyInput
	}

	result, err := service.adminRepository.Login(admin.Username)
	if err != nil {
		return &web.AdminLoginResponse{}, constant.ErrLoginFailed
	}

	err = util.CheckPassword(admin.Password, result.Password)
	if err != nil {
		return &web.AdminLoginResponse{}, constant.ErrLoginFailed
	}

	token, err := middleware.CreateToken(int(result.ID), result.Username, false)

	response := web.AdminLoginResponse{
		Name:     result.Name,
		Username: result.Username,
		Token:    token,
	}

	return &response, nil
}
