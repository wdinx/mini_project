package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util/converter"
)

type TicketServiceImpl struct {
	ticketRepository repository.TicketRepository
	validator        *validator.Validate
}

func NewTicketService(ticketRepository repository.TicketRepository, validator *validator.Validate) TicketService {
	return &TicketServiceImpl{ticketRepository: ticketRepository, validator: validator}
}

func (service *TicketServiceImpl) FindByID(id string) (*web.TicketResponse, error) {
	ticket, err := service.ticketRepository.FindByID(id)

	response := converter.ToTicketResponse(ticket)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *TicketServiceImpl) FindByUserID(userID int) (*[]web.TicketResponse, error) {
	tickets, err := service.ticketRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var response []web.TicketResponse
	for _, ticket := range *tickets {
		response = append(response, *converter.ToTicketResponse(&ticket))
	}
	return &response, nil
}

func (service *TicketServiceImpl) FindByTouristAttractionID(touristAttractionID int) (*[]web.TicketResponse, error) {
	tickets, err := service.ticketRepository.FindByTouristAttractionID(touristAttractionID)
	if err != nil {
		return nil, err
	}
	var response []web.TicketResponse
	for _, ticket := range *tickets {
		response = append(response, *converter.ToTicketResponse(&ticket))
	}
	return &response, nil
}
