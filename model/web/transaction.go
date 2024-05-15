package web

import (
	"github.com/google/uuid"
	"time"
)

type TransactionCreateRequest struct {
	UserID              uint   `json:"user_id" validate:"required"`
	TouristAttractionID int    `json:"tourist_attraction_id" validate:"required"`
	Qty                 int    `json:"qty"`
	ReservationDate     string `json:"reservation_date" validate:"required"`
	Status              string `json:"status"`
}

type TransactionCreateResponse struct {
	User                      UserLoginResponse         `json:"user"`
	TouristAttractionResponse TouristAttractionResponse `json:"tourist_attraction"`
	Qty                       int                       `json:"qty"`
	Amount                    float64                   `json:"amount"`
}

type TransactionResponse struct {
	ID                uuid.UUID                 `json:"id"`
	TouristAttraction TouristAttractionResponse `json:"tourist_attraction"`
	Qty               int                       `json:"qty"`
	Amount            float64                   `json:"amount"`
	ReservationDate   time.Time                 `json:"reservation_date"`
	Status            string                    `json:"status"`
}

type UserTransactionResponse struct {
	ID                  uint                  `json:"id"`
	Name                string                `json:"name"`
	Email               string                `json:"email"`
	NoPhone             string                `json:"no_phone"`
	ProfilePicture      string                `json:"profile_picture"`
	TransactionResponse []TransactionResponse `json:"transactions"`
}
