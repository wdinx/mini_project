package web

import "mime/multipart"

type UserRegisterRequest struct {
	Name           string                `json:"name" form:"name" validate:"required,min=1,max=200"`
	Email          string                `json:"email" form:"email" validate:"required,email,min=1,max=200"`
	NoPhone        string                `json:"no_phone" form:"no_phone" validate:"required,min=1,max=200"`
	Password       string                `json:"password" form:"password" validate:"required,min=6,max=200"`
	ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email,min=1,max=200"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=200"`
}

type UserLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	Image string `json:"image"`
}
