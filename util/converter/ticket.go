package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
)

func ToTicketResponse(ticket *domain.Ticket) *web.TicketResponse {
	return &web.TicketResponse{
		Id:                ticket.ID,
		ReservationDate:   ticket.ReservationDate,
		TouristAttraction: *ToTouristAttractionResponse(&ticket.TouristAttraction),
	}
}

func ToUserTicketResponse(user *domain.User, tickets *[]domain.Ticket) *web.UserTicketResponse {
	var response []web.TicketResponse
	for _, value := range *tickets {
		response = append(response, *ToTicketResponse(&value))
	}
	return &web.UserTicketResponse{
		Id:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		NoPhone:        user.NoPhone,
		ProfilePicture: user.ProfilePicture,
		TicketResponse: response,
	}
}
