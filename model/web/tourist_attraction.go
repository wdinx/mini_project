package web

import (
	"mime/multipart"
	"mini_project/model/domain"
)

type TouristAttractionRequest struct {
	Name                    string                `json:"name" form:"name" validate:"required,min=1,max=200"`
	Description             string                `json:"description" form:"description" validate:"required,min=1,max=200"`
	TouristAttractionTypeId int                   `json:"tourist_attraction_type_id" form:"tourist_attraction_type_id" validate:"required"`
	Location                string                `json:"location" form:"location" validate:"required,min=1,max=200"`
	TicketPrice             string                `json:"ticket_price" form:"ticket_price" validate:"required,min=1,max=200"`
	Image                   *multipart.FileHeader `json:"image" form:"image" validate:"required"`
}

type TouristAttractionUpdateRequest struct {
	ID                      int                   `json:"id" form:"id"`
	Name                    string                `json:"name" form:"name" validate:",min=1,max=200"`
	Description             string                `json:"description" form:"description" validate:",min=1,max=200"`
	TouristAttractionTypeId int                   `json:"tourist_attraction_type_id" form:"tourist_attraction_type_id"`
	Location                string                `json:"location" form:"location" validate:",min=1,max=200"`
	TicketPrice             string                `json:"ticket_price" form:"ticket_price" validate:",min=1,max=200"`
	Image                   *multipart.FileHeader `json:"image" form:"image"`
	Balance                 int                   `json:"balance" form:"balance"`
}

type TouristAttractionResponse struct {
	Id                    int                          `json:"id"`
	Name                  string                       `json:"name"`
	Description           string                       `json:"description"`
	TouristAttractionType domain.TouristAttractionType `json:"tourist_attraction_type"`
	Location              string                       `json:"location"`
	TicketPrice           string                       `json:"ticket_price"`
	Image                 string                       `json:"image"`
	Balance               int                          `json:"balance"`
}

type TouristAttractionBalanceResponse struct {
	Id      int `json:"id"`
	Balance int `json:"balance"`
}