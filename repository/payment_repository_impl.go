package repository

import (
	"gorm.io/gorm"
	"mini_project/constant"
	"mini_project/model/domain"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{db}
}

func (repository *PaymentRepositoryImpl) FindByID(id string) (payment *domain.Payment, err error) {
	if err = repository.db.Preload("Transaction").First(&payment, "id LIKE ?", id).Error; err != nil {
		return payment, constant.ErrDataNotFound
	}
	return payment, nil
}

func (repository *PaymentRepositoryImpl) Insert(payment *domain.Payment) error {
	if err := repository.db.Create(&payment).Error; err != nil {
		return constant.ErrInsertData
	}
	return nil
}

func (repository *PaymentRepositoryImpl) Update(payment *domain.Payment) error {
	if err := repository.db.Save(&payment).Error; err != nil {
		return constant.ErrUpdateData
	}
	return nil
}
