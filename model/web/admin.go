package web

type AdminRegisterRequest struct {
	Name                string `json:"name" validate:"required,min=1,max=200"`
	Username            string `json:"username" validate:"required,min=1,max=200"`
	Password            string `json:"password" validate:"required,min=6,max=200"`
	TouristAttractionID int    `json:"tourist_attraction_id" validate:"required"`
}

type AdminRegisterResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
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
