package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"mini_project/middleware"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	validator      *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{userRepository: userRepository, validator: validator}
}

func (service *UserServiceImpl) Login(request web.UserLoginRequest) (*web.UserLoginResponse, error) {
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

	response := web.UserLoginResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Image: user.ProfilePicture,
		Token: token,
	}

	return &response, nil
}

func (service *UserServiceImpl) Register(request web.UserRegisterRequest) (domain.User, error) {
	var err error
	err = service.validator.Struct(request)
	if err != nil {
		return domain.User{}, err
	}

	request.Password, err = util.HashPassword(request.Password)
	util.PanicIfError(err)

	newFileName := util.StoreImageToLocal(request.ProfilePicture, request.Name)
	fmt.Println(newFileName)

	user := domain.User{
		Name:           request.Name,
		Email:          request.Email,
		Password:       request.Password,
		NoPhone:        request.NoPhone,
		ProfilePicture: util.GetImageUrl(newFileName),
	}

	err = service.userRepository.Register(&user)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
