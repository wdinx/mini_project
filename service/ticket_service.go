package service

import (
	"mini_project/model/web"
)

type TicketService interface {
	FindByID(id string) (*web.TicketResponse, error)
	FindByUserID(userID int) (*[]web.TicketResponse, error)
	FindByTouristAttractionID(touristAttractionID int) (*[]web.TicketResponse, error)
}
