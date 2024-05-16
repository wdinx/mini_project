package service

import (
	"mini_project/model/web"
)

type TransactionService interface {
	Create(request *web.TransactionCreateRequest) (*web.PaymentResponse, error)
	Delete(transactionID int) error
	GetByUserID(userID int) (*web.UserTransactionResponse, error)
	GetByID(transactionID string) (*web.TransactionResponse, error)
}
