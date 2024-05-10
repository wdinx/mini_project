package _interface

import "mini_project/model/domain"

type MidtransService interface {
	GenerateSnapURL(payment *domain.Payment) error
	VerifyPayment(orderID string) (bool, error)
}
