package web

import (
	"github.com/google/uuid"
	"time"
)

type TicketResponse struct {
	Id                uuid.UUID                 `json:"id"`
	ReservationDate   time.Time                 `json:"reservation_date"`
	TouristAttraction TouristAttractionResponse `json:"tourist_attraction"`
}

type UserTicketResponse struct {
	Id             uint             `json:"id"`
	Name           string           `json:"name"`
	Email          string           `json:"email"`
	NoPhone        string           `json:"no_phone"`
	ProfilePicture string           `json:"profile_picture"`
	TicketResponse []TicketResponse `json:"tickets"`
}
