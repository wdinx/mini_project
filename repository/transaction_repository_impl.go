package repository

import (
	"gorm.io/gorm"
	"mini_project/model/domain"
	"mini_project/repository/interface"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) _interface.TransactionRepository {
	return &TransactionRepositoryImpl{DB: DB}
}

func (repository *TransactionRepositoryImpl) Create(transaction *domain.Transaction) (*domain.Transaction, error) {
	if err := repository.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (repository *TransactionRepositoryImpl) Delete(transactionID int) error {
	if err := repository.DB.Unscoped().Delete(&domain.Transaction{}, "id LIKE ?", transactionID).Error; err != nil {
		return err
	}
	return nil
}
