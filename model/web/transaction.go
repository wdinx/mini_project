package web

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
