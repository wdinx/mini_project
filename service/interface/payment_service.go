package _interface

import "mini_project/model/web"

type PaymentService interface {
	ConfirmedPayment(id string) error
	InitializePayment(request *web.PaymentRequest) (*web.PaymentResponse, error)
}
