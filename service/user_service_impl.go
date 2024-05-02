package service

import (
	"fmt"
	"mini_project/middleware"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (service UserServiceImpl) Login(request web.UserLoginRequest) (*web.UserLoginResponse, error) {
	user, err := service.userRepository.Login(request.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(request.Password, user.Password)
	if err != nil {
		return nil, err
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name, true)
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

func (service UserServiceImpl) Register(request web.UserRegisterRequest) (domain.User, error) {
	var err error
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
