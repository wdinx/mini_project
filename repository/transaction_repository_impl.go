package repository

import (
	"gorm.io/gorm"
	"mini_project/model/domain"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
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

func (repository *TransactionRepositoryImpl) GetByUserID(userID int) (transactions *[]domain.Transaction, err error) {
	if err = repository.DB.Preload("TouristAttraction").Preload("TouristAttraction.TouristAttractionType").Preload("User").Find(&transactions, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (repository *TransactionRepositoryImpl) GetByID(transactionID int) (transaction *domain.Transaction, err error) {
	if err = repository.DB.Preload("TouristAttraction").Preload("TouristAttractionType").Preload("User").Find(&transaction, "id = ?", transactionID).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
