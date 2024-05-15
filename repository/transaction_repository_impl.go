package repository

import (
	"gorm.io/gorm"
	"mini_project/constant"
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
		return &domain.Transaction{}, constant.ErrInsertData
	}
	return transaction, nil
}

func (repository *TransactionRepositoryImpl) Delete(transactionID int) error {
	if err := repository.DB.Unscoped().Delete(&domain.Transaction{}, "id LIKE ?", transactionID).Error; err != nil {
		return constant.ErrDeleteData
	}
	return nil
}

func (repository *TransactionRepositoryImpl) GetByUserID(userID int) (transactions *[]domain.Transaction, err error) {
	if err = repository.DB.Preload("TouristAttraction").Preload("TouristAttraction.TouristAttractionType").Find(&transactions, "user_id = ?", userID).Error; err != nil {
		return transactions, constant.ErrDataNotFound
	}
	return transactions, nil
}

func (repository *TransactionRepositoryImpl) GetByID(transactionID int) (transaction *domain.Transaction, err error) {
	if err = repository.DB.Preload("TouristAttraction").Preload("TouristAttractionType").Preload("User").Find(&transaction, "id = ?", transactionID).Error; err != nil {
		return transaction, constant.ErrDataNotFound
	}
	return transaction, nil
}

func (repository *TransactionRepositoryImpl) Update(transaction *domain.Transaction) error {
	if err := repository.DB.Save(&transaction).Error; err != nil {
		return constant.ErrUpdateData
	}
	return nil
}
