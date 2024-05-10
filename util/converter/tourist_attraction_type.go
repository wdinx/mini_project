package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
)

func ToCreateTouristAttractionType(request *web.TouristAttractionTypeCreateRequest) *domain.TouristAttractionType {
	return &domain.TouristAttractionType{
		Name: request.Name,
	}
}

func ToTouristAttractionTypeResponse(touristAttractionType *domain.TouristAttractionType) *web.TouristAttractionTypeResponse {
	return &web.TouristAttractionTypeResponse{
		ID:   touristAttractionType.ID,
		Name: touristAttractionType.Name,
	}
}
