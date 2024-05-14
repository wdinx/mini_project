package repository

import (
	"gorm.io/gorm"
	"mini_project/constant"
	"mini_project/model/domain"
)

type TouristAttractionTypeRepositoryImpl struct {
	DB *gorm.DB
}

func NewTouristAttractionTypeRepository(DB *gorm.DB) TouristAttractionTypeRepository {
	return &TouristAttractionTypeRepositoryImpl{DB}
}

func (repository *TouristAttractionTypeRepositoryImpl) Create(touristAttractionType *domain.TouristAttractionType) (*domain.TouristAttractionType, error) {
	if err := repository.DB.Create(&touristAttractionType).Error; err != nil {
		return &domain.TouristAttractionType{}, constant.ErrInsertData
	}
	return touristAttractionType, nil
}

func (repository *TouristAttractionTypeRepositoryImpl) Update(touristAttractionType *domain.TouristAttractionType) (*domain.TouristAttractionType, error) {
	if err := repository.DB.Save(&touristAttractionType).Where("id = ?", touristAttractionType.ID).Error; err != nil {
		return &domain.TouristAttractionType{}, err
	}
	return touristAttractionType, nil
}

func (repository *TouristAttractionTypeRepositoryImpl) Delete(touristAttractionTypeId int) error {

	var touristAttractionType *domain.TouristAttractionType
	if err := repository.DB.Take(&touristAttractionType, "id = ?", touristAttractionTypeId).Error; err != nil {
		return err
	}

	if err := repository.DB.Delete(touristAttractionType).Error; err != nil {
		return err
	}
	return nil
}

func (repository *TouristAttractionTypeRepositoryImpl) GetAll() []domain.TouristAttractionType {
	var touristAttractionTypes []domain.TouristAttractionType
	if err := repository.DB.Find(&touristAttractionTypes).Error; err != nil {
		return []domain.TouristAttractionType{}
	}
	return touristAttractionTypes
}

func (repository *TouristAttractionTypeRepositoryImpl) FindByID(id int) (*domain.TouristAttractionType, error) {
	var touristAttractionType domain.TouristAttractionType
	if err := repository.DB.Find(&touristAttractionType, "id = ?", id).Error; err != nil {
		return &domain.TouristAttractionType{}, err
	}
	return &touristAttractionType, nil
}
