package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/util"
)

func ToUserModel(request *web.UserRegisterRequest, filename string) *domain.User {

	var err error
	request.Password, err = util.HashPassword(request.Password)
	util.PanicIfError(err)

	return &domain.User{
		Name:           request.Name,
		Email:          request.Email,
		NoPhone:        request.NoPhone,
		Password:       request.Password,
		ProfilePicture: filename,
	}
}

func ToUserLoginResponse(admin *domain.User, token string) *web.UserLoginResponse {
	return &web.UserLoginResponse{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Token: token,
	}
}

func ToUserResponse(user *domain.User) *web.UserResponse {
	return &web.UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		NoPhone:        user.NoPhone,
		ProfilePicture: util.GetImageUrl(user.ProfilePicture),
	}
}
