package service

import "mini_project/model/web"

type TouristAttractionTypeService interface {
	Create(request *web.TouristAttractionTypeCreateRequest) (*web.TouristAttractionTypeResponse, error)
	Update(request *web.TouristAttractionTypeUpdateRequest) (*web.TouristAttractionTypeResponse, error)
	Delete(touristAttractionTypeId int) error
	GetAll() *[]web.TouristAttractionTypeResponse
	FindByID(touristAttractionTypeId int) (*web.TouristAttractionTypeResponse, error)
}
