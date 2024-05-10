package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/constant"
	"mini_project/middleware"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util"
	"mini_project/util/converter"
)

type AdminServiceImpl struct {
	adminRepository repository.AdminRepository
	validator       *validator.Validate
}

func NewAdminRepository(adminRepository repository.AdminRepository, validator *validator.Validate) AdminService {
	return &AdminServiceImpl{
		adminRepository: adminRepository,
		validator:       validator,
	}
}

func (service *AdminServiceImpl) Register(admin *web.AdminRegisterRequest) (*web.AdminRegisterResponse, error) {
	err := service.validator.Struct(admin)
	if err != nil {
		return &web.AdminRegisterResponse{}, err
	}

	admin.Password, err = util.HashPassword(admin.Password)
	util.PanicIfError(err)

	insertAdmin := converter.ToAdminModel(admin)

	adminResponse, err := service.adminRepository.Register(insertAdmin)
	if err != nil {
		return &web.AdminRegisterResponse{}, constant.ErrInsertDatabase
	}

	result := converter.ToAdminRegisterResponse(adminResponse)

	return result, nil
}

func (service *AdminServiceImpl) Login(admin web.AdminLoginRequest) (*web.AdminLoginResponse, error) {
	if err := service.validator.Struct(admin); err != nil {
		return &web.AdminLoginResponse{}, err
	}

	result, err := service.adminRepository.Login(admin.Username)
	if err != nil {
		return &web.AdminLoginResponse{}, constant.ErrLoginFailed
	}

	err = util.CheckPassword(admin.Password, result.Password)
	if err != nil {
		return &web.AdminLoginResponse{}, constant.ErrLoginFailed
	}

	token, err := middleware.CreateTokeForAdmin(int(result.ID), result.Username)

	response := converter.ToAdminLoginResponse(result, token)

	return response, nil
}
