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

func ToTransactionResponse(transaction *domain.Transaction) *web.TransactionResponse {
	return &web.TransactionResponse{
		ID:                transaction.ID,
		TouristAttraction: *ToTouristAttractionResponse(&transaction.TouristAttraction),
		Qty:               transaction.Qty,
		Amount:            transaction.Amount,
		ReservationDate:   transaction.ReservationDate,
		Status:            transaction.Status,
	}
}

func ToUserTransactionResponse(user *domain.User, transactions *[]domain.Transaction) *web.UserTransactionResponse {
	var response []web.TransactionResponse
	for _, value := range *transactions {
		response = append(response, *ToTransactionResponse(&value))
	}
	return &web.UserTransactionResponse{
		ID:                  user.ID,
		Name:                user.Name,
		Email:               user.Email,
		NoPhone:             user.NoPhone,
		ProfilePicture:      user.ProfilePicture,
		TransactionResponse: response,
	}
}
