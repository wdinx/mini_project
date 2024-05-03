package web

type AdminRegisterRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=200"`
	Username string `json:"username" validate:"required,min=1,max=200"`
	Password string `json:"password" validate:"required,min=6,max=200"`
}

type AdminRegisterResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required,min=1,max=200"`
	Password string `json:"password" validate:"required,min=6,max=200"`
}
