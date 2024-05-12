package converter

import (
	"github.com/google/uuid"
	"mini_project/model/domain"
	"mini_project/model/web"
	"time"
)

func ToTransactionModel(request *web.TransactionCreateRequest, touristAttraction *domain.TouristAttraction, reservationDate time.Time) *domain.Transaction {

	return &domain.Transaction{
		ID:                  uuid.New(),
		UserID:              request.UserID,
		TouristAttractionID: request.TouristAttractionID,
		Qty:                 request.Qty,
		Amount:              touristAttraction.TicketPrice * float64(request.Qty),
		ReservationDate:     reservationDate,
		Status:              "PENDING",
	}
}
