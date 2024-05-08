package service

import (
	"mini_project/model/web"
)

type TransactionService interface {
	Create(request *web.TransactionCreateRequest) error
	Delete(transactionID int) error
}