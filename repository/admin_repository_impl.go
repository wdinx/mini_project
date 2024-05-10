package repository

import (
	"gorm.io/gorm"
	"mini_project/model/domain"
	"mini_project/repository/interface"
)

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) _interface.AdminRepository {
	return &AdminRepositoryImpl{
		DB: db,
	}
}

func (repository *AdminRepositoryImpl) Register(admin *domain.Admin) (*domain.Admin, error) {
	if err := repository.DB.Create(&admin).Error; err != nil {
		return &domain.Admin{}, err
	}
	return admin, nil
}

func (repository *AdminRepositoryImpl) Login(username string) (*domain.Admin, error) {
	var admin domain.Admin
	if err := repository.DB.First(&admin, "username LIKE ?", username).Limit(1).Error; err != nil {
		return &admin, err
	}
	return &admin, nil
}

func (repository *AdminRepositoryImpl) GetByID(id int) (*domain.Admin, error) {
	var admin domain.Admin
	if err := repository.DB.Preload("TouristAttraction").First(&admin, id).Error; err != nil {
		return &admin, err
	}
	return &admin, nil
}
