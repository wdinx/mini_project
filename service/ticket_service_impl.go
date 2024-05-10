package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/domain"
	_interface2 "mini_project/repository/interface"
	"mini_project/service/interface"
)

type TicketServiceImpl struct {
	ticketRepository _interface2.TicketRepository
	validator        *validator.Validate
}

func NewTicketService(ticketRepository _interface2.TicketRepository, validator *validator.Validate) _interface.TicketService {
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
