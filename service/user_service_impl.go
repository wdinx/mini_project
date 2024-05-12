package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/middleware"
	"mini_project/model/domain"
	"mini_project/model/web"
	_interface2 "mini_project/repository"
	"mini_project/util"
	"mini_project/util/converter"
)

type UserServiceImpl struct {
	userRepository _interface2.UserRepository
	validator      *validator.Validate
}

func NewUserService(userRepository _interface2.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{userRepository: userRepository, validator: validator}
}

func (service *UserServiceImpl) Login(request *web.UserLoginRequest) (*web.UserLoginResponse, error) {
	if err := service.validator.Struct(request); err != nil {
		return nil, err
	}
	user, err := service.userRepository.Login(request.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(request.Password, user.Password)
	if err != nil {
		return nil, err
	}

	token, err := middleware.CreateTokenForUser(int(user.ID), user.Name)
	util.PanicIfError(err)

	response := converter.ToUserLoginResponse(user, token)

	return response, nil
}

func (service *UserServiceImpl) Register(request *web.UserRegisterRequest) (*domain.User, error) {
	var err error
	err = service.validator.Struct(request)
	if err != nil {
		return nil, err
	}

	user := converter.ToUserModel(request)

	err = service.userRepository.Register(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
