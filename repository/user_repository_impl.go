package repository

import (
	"gorm.io/gorm"
	"mini_project/model/domain"
	"mini_project/repository/interface"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) _interface.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Login(email string) (*domain.User, error) {
	var user domain.User
	if err := repository.DB.First(&user, "email LIKE ?", email).Error; err != nil {
		return &domain.User{}, err
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) Register(user *domain.User) error {
	if err := repository.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
