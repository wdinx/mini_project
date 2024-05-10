package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/constant"
	"mini_project/middleware"
	"mini_project/model/domain"
	"mini_project/model/web"
	_interface2 "mini_project/repository/interface"
	"mini_project/service/interface"
	"mini_project/util"
)

type AdminServiceImpl struct {
	adminRepository _interface2.AdminRepository
	validator       *validator.Validate
}

func NewAdminRepository(adminRepository _interface2.AdminRepository, validator *validator.Validate) _interface.AdminService {
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

	insertAdmin := domain.Admin{
		Name:                admin.Name,
		Username:            admin.Username,
		Password:            admin.Password,
		TouristAttractionID: admin.TouristAttractionID,
	}

	adminResponse, err := service.adminRepository.Register(&insertAdmin)
	if err != nil {
		return &web.AdminRegisterResponse{}, constant.ErrInsertDatabase
	}

	result := web.AdminRegisterResponse{
		Name:     adminResponse.Name,
		Username: adminResponse.Username,
	}

	return &result, nil
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

	response := web.AdminLoginResponse{
		Name:     result.Name,
		Username: result.Username,
		Token:    token,
	}

	return &response, nil
}
