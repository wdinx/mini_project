package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
)

func ToTicketResponse(ticket *domain.Ticket) *web.TicketResponse {
	return &web.TicketResponse{
		Id:                ticket.ID,
		ReservationDate:   ticket.ReservationDate,
		User:              *ToUserResponse(&ticket.User),
		TouristAttraction: *ToTouristAttractionResponse(&ticket.TouristAttraction),
	}
}
