package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/util"
)

func ToUserModel(request *web.UserRegisterRequest) *domain.User {

	var err error
	request.Password, err = util.HashPassword(request.Password)
	util.PanicIfError(err)

	newFileName := util.StoreImageToLocal(request.ProfilePicture, request.Name)

	return &domain.User{
		Name:           request.Name,
		Email:          request.Email,
		NoPhone:        request.NoPhone,
		Password:       request.Password,
		ProfilePicture: newFileName,
	}
}

func ToUserRegisterResponse(admin *domain.Admin) *web.AdminRegisterResponse {
	return &web.AdminRegisterResponse{
		Name:     admin.Name,
		Username: admin.Username,
	}
}

func ToUserLoginResponse(admin *domain.User, token string) *web.UserResponse {
	return &web.UserResponse{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Token: token,
	}
}
