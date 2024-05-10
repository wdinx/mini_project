package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/domain"
	"mini_project/repository"
)

type TicketServiceImpl struct {
	ticketRepository repository.TicketRepository
	validator        *validator.Validate
}

func NewTicketService(ticketRepository repository.TicketRepository, validator *validator.Validate) TicketService {
	return &TicketServiceImpl{ticketRepository: ticketRepository, validator: validator}
}

func (service *TicketServiceImpl) FindByID(id string) (*domain.Ticket, error) {
	ticket, err := service.ticketRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (service *TicketServiceImpl) FindByUserID(userID int) (*[]domain.Ticket, error) {
	tickets, err := service.ticketRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
