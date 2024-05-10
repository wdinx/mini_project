package web

import "github.com/google/uuid"

type PaymentRequest struct {
	Amount        float64   `validate:"required" json:"amount" form:"amount"`
	TransactionID uuid.UUID `validate:"required" json:"transaction_id" form:"transaction_id"`
}

type PaymentResponse struct {
	ID      uuid.UUID `json:"id"`
	Amount  float64   `json:"amount"`
	SnapURL string    `json:"snap_url"`
	Status  int       `json:"status"`
}
