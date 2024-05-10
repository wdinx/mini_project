package repository

import "mini_project/model/domain"

type TicketRepository interface {
	FindByID(id string) (*domain.Ticket, error)
	FindByUserID(userID int) (*[]domain.Ticket, error)
	FindByTouristAttractionID(touristAttractionID int) (*[]domain.Ticket, error)
	Insert(ticket *domain.Ticket) error
}
