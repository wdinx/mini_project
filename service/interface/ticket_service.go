package _interface

import "mini_project/model/domain"

type TicketService interface {
	FindByID(id string) (*domain.Ticket, error)
	FindByUserID(userID int) (*[]domain.Ticket, error)
}
