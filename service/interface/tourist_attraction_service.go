package _interface

import (
	"mini_project/model/web"
)

type TouristAttractionService interface {
	Create(request *web.TouristAttractionRequest) (*web.TouristAttractionResponse, error)
	Update(request *web.TouristAttractionUpdateRequest) (*web.TouristAttractionResponse, error)
	UpdateBalanceById(request *web.TouristAttractionUpdateRequest) error
	GetAllTouristAttraction() (*[]web.TouristAttractionResponse, error)
}
