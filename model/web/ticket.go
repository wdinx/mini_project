package web

import (
	"github.com/google/uuid"
	"time"
)

type TicketResponse struct {
	Id                uuid.UUID                 `json:"id"`
	ReservationDate   time.Time                 `json:"reservation_date"`
	User              UserResponse              `json:"user"`
	TouristAttraction TouristAttractionResponse `json:"tourist_attraction"`
}
