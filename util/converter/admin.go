package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
)

func ToAdminModel(request *web.AdminRegisterRequest) *domain.Admin {
	return &domain.Admin{
		Username:            request.Username,
		Password:            request.Password,
		Name:                request.Name,
		TouristAttractionID: request.TouristAttractionID,
	}
}

func ToAdminResponse(admin *domain.Admin) *web.AdminRegisterResponse {
	return &web.AdminRegisterResponse{
		Name:     admin.Name,
		Username: admin.Username,
	}
}
