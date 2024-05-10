package repository

import "mini_project/model/domain"

type TouristAttractionRepository interface {
	Create(touristAttraction *domain.TouristAttraction) (*domain.TouristAttraction, error)
	Update(touristAttraction *domain.TouristAttraction) (*domain.TouristAttraction, error)
	UpdateBalanceById(touristAttraction *domain.TouristAttraction) (*domain.TouristAttraction, error)
	GetAllTouristAttraction() (*[]domain.TouristAttraction, error)
	FindByID(touristAttractionId int) (*domain.TouristAttraction, error)
	Delete(touristAttractionId int) error
}
