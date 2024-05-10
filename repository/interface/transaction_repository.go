package _interface

import "mini_project/model/domain"

type TransactionRepository interface {
	Create(transaction *domain.Transaction) (*domain.Transaction, error)
	Delete(transactionID int) error
}
