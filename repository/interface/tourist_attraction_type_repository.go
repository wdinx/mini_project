package _interface

import "mini_project/model/domain"

type TouristAttractionTypeRepository interface {
	Create(touristAttractionType *domain.TouristAttractionType) (*domain.TouristAttractionType, error)
	Update(touristAttractionType *domain.TouristAttractionType) (*domain.TouristAttractionType, error)
	Delete(touristAttractionTypeId int) error
	FindByID(id int) (*domain.TouristAttractionType, error)
	GetAll() []domain.TouristAttractionType
}
