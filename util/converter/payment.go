package converter

import (
	"github.com/google/uuid"
	"mini_project/model/domain"
	"mini_project/model/web"
)

func ToPaymentModel(request *web.PaymentRequest) *domain.Payment {
	return &domain.Payment{
		ID:            uuid.New(),
		TransactionID: request.TransactionID,
		Amount:        request.Amount,
	}
}

func ToPaymentResponse(payment *domain.Payment) *web.PaymentResponse {
	return &web.PaymentResponse{
		ID:      payment.ID,
		Amount:  payment.Amount,
		SnapURL: payment.SnapURL,
		Status:  payment.Status,
	}
}
