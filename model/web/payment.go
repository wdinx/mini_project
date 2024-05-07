package web

type PaymentRequest struct {
	Amount              float64 `validate:"required" json:"amount" form:"amount"`
	UserID              int     `validate:"required" json:"user_id" form:"user_id"`
	TouristAttractionID int     `validate:"required" json:"tourist_attraction_id" form:"tourist_attraction_id"`
}

type PaymentResponse struct {
	ID                string                    `json:"id"`
	Amount            float64                   `json:"amount"`
	User              UserLoginResponse         `json:"user"`
	TouristAttraction TouristAttractionResponse `json:"tourist_attraction"`
	SnapURL           string                    `json:"snap_url"`
}
