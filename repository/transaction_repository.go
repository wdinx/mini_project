package repository

import "mini_project/model/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) (*domain.Transaction, error)
	Delete(transactionID int) error
	GetByUserID(userID int) (*[]domain.Transaction, error)
	GetByID(transactionID int) (*domain.Transaction, error)
}
