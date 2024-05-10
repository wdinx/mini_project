package repository

import (
	"gorm.io/gorm"
	"mini_project/model/domain"
)

type TouristAttractionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTouristAttractionRepository(db *gorm.DB) TouristAttractionRepository {
	return &TouristAttractionRepositoryImpl{DB: db}
}

func (repository *TouristAttractionRepositoryImpl) Create(touristAttraction *domain.TouristAttraction) (*domain.TouristAttraction, error) {
	if err := repository.DB.Create(touristAttraction).Error; err != nil {
		return nil, err
	}
	return touristAttraction, nil
}

func (repository *TouristAttractionRepositoryImpl) UpdateBalanceById(touristAttraction *domain.TouristAttraction) (*domain.TouristAttraction, error) {
	if err := repository.DB.Model(&domain.TouristAttraction{}).Where("id = ?", touristAttraction.ID).Update("balance", touristAttraction.Balance).Error; err != nil {
		return nil, err
	}
	return touristAttraction, nil
}

func (repository *TouristAttractionRepositoryImpl) Delete(touristAttractionId int) error {
	var touristAttraction *domain.TouristAttraction
	if err := repository.DB.Take(&touristAttraction, "id = ?", touristAttractionId).Error; err != nil {
		return err
	}

	if err := repository.DB.Delete(touristAttraction).Error; err != nil {
		return err
	}

	return nil
}

func (repository *TouristAttractionRepositoryImpl) Update(touristAttraction *domain.TouristAttraction) (*domain.TouristAttraction, error) {
	if err := repository.DB.Save(touristAttraction).Where("id = ?", touristAttraction.ID).Error; err != nil {
		return nil, err
	}
	return touristAttraction, nil
}

func (repository *TouristAttractionRepositoryImpl) GetAllTouristAttraction() (*[]domain.TouristAttraction, error) {
	var touristAttractions *[]domain.TouristAttraction
	if err := repository.DB.Model(&domain.TouristAttraction{}).Joins("TouristAttractionType").Find(&touristAttractions).Error; err != nil {
		return nil, err
	}
	return touristAttractions, nil
}

func (repository *TouristAttractionRepositoryImpl) FindByID(touristAttractionId int) (*domain.TouristAttraction, error) {
	var touristAttraction *domain.TouristAttraction
	if err := repository.DB.First(&touristAttraction, "id LIKE ?", touristAttractionId).Error; err != nil {
		return nil, err
	}
	return touristAttraction, nil
}
