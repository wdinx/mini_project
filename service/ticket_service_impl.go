package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util/converter"
)

type TicketServiceImpl struct {
	ticketRepository repository.TicketRepository
	userRepository   repository.UserRepository
	validator        *validator.Validate
}

func NewTicketService(ticketRepository repository.TicketRepository, userRepository repository.UserRepository, validator *validator.Validate) TicketService {
	return &TicketServiceImpl{ticketRepository: ticketRepository, userRepository: userRepository, validator: validator}
}

func (service *TicketServiceImpl) FindByID(id string) (*web.TicketResponse, error) {
	ticket, err := service.ticketRepository.FindByID(id)
	if err != nil {
		return &web.TicketResponse{}, err
	}

	response := converter.ToTicketResponse(ticket)

	return response, nil
}

func (service *TicketServiceImpl) FindByUserID(userID int) (*web.UserTicketResponse, error) {
	tickets, err := service.ticketRepository.FindByUserID(userID)
	if err != nil {
		return &web.UserTicketResponse{}, err
	}

	user, err := service.userRepository.GetByID(userID)
	if err != nil {
		return &web.UserTicketResponse{}, err
	}

	response := converter.ToUserTicketResponse(user, tickets)

	return response, nil
}

func (service *TicketServiceImpl) FindByTouristAttractionID(touristAttractionID int) (*[]web.TicketResponse, error) {
	tickets, err := service.ticketRepository.FindByTouristAttractionID(touristAttractionID)
	if err != nil {
		return &[]web.TicketResponse{}, err
	}
	var response []web.TicketResponse
	for _, ticket := range *tickets {
		response = append(response, *converter.ToTicketResponse(&ticket))
	}
	return &response, nil
}
