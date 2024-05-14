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

type UserServiceImpl struct {
	userRepository repository.UserRepository
	imageService   ImageService
	validator      *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, imageService ImageService, validator *validator.Validate) UserService {
	return &UserServiceImpl{userRepository: userRepository, imageService: imageService, validator: validator}
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
		return nil, constant.ErrLogin
	}

	token, err := middleware.CreateTokenForUser(int(user.ID), user.Name)
	util.PanicIfError(err)

	response := converter.ToUserLoginResponse(user, token)

	return response, nil
}

func (service *UserServiceImpl) Register(request *web.UserRegisterRequest) (*web.UserResponse, error) {
	var err error
	err = service.validator.Struct(request)
	if err != nil {
		return &web.UserResponse{}, constant.ErrEmptyInput
	}

	filename := util.GenerateImageName(request.Name, request.ProfilePicture.Filename)

	err = service.imageService.UploadImage(request.ProfilePicture, util.GenerateImageName(request.Name, filename))
	if err != nil {
		return &web.UserResponse{}, err
	}

	user := converter.ToUserModel(request, filename)

	err = service.userRepository.Register(user)

	if err != nil {
		return &web.UserResponse{}, err
	}

	response := converter.ToUserResponse(user)

	return response, nil
}
