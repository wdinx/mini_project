package converter

import (
	"github.com/google/uuid"
	"mini_project/model/domain"
	"mini_project/model/web"
)

func ToTransactionModel(request *web.TransactionCreateRequest, touristAttraction *domain.TouristAttraction) *domain.Transaction {
	return &domain.Transaction{
		ID:                  uuid.New(),
		UserID:              request.UserID,
		TouristAttractionID: request.TouristAttractionID,
		Qty:                 request.Qty,
		Amount:              touristAttraction.TicketPrice * float64(request.Qty),
		ReservationDate:     request.ReservationDate,
		Status:              "PENDING",
	}
}
