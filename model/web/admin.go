package web

type AdminRegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
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
	Username string `json:"username"`
	Password string `json:"password"`
}
