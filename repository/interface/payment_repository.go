package _interface

import "mini_project/model/domain"

type PaymentRepository interface {
	FindByID(id string) (*domain.Payment, error)
	Insert(payment *domain.Payment) error
	Update(payment *domain.Payment) error
}
