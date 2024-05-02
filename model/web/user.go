package web

import "mime/multipart"

type UserRegisterRequest struct {
	Name           string                `json:"name" form:"name"`
	Email          string                `json:"email" form:"email"`
	NoPhone        string                `json:"no_phone" form:"no_phone"`
	Password       string                `json:"password" form:"password"`
	ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
	Image string `json:"image"`
}
