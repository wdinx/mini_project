package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"mini_project/model/domain"
	"mini_project/model/web"
	_interface2 "mini_project/repository/interface"
	"mini_project/service/interface"
	"time"
)

type TransactionServiceImpl struct {
	transactionRepository       _interface2.TransactionRepository
	touristAttractionRepository _interface2.TouristAttractionRepository
	paymentService              _interface.PaymentService
	validator                   *validator.Validate
}

func NewTransactionService(
	transactionRepository _interface2.TransactionRepository,
	touristAttractionRepository _interface2.TouristAttractionRepository,
	paymentService _interface.PaymentService,
	validator *validator.Validate) _interface.TransactionService {
	return &TransactionServiceImpl{
		transactionRepository:       transactionRepository,
		touristAttractionRepository: touristAttractionRepository,
		paymentService:              paymentService,
		validator:                   validator,
	}
}

func (service *TransactionServiceImpl) Create(request *web.TransactionCreateRequest) (*web.PaymentResponse, error) {
	if err := service.validator.Struct(request); err != nil {
		return nil, err
	}

	touristAttraction, err := service.touristAttractionRepository.FindByID(request.TouristAttractionID)
	if err != nil {
		return nil, err
	}

	reservationDate, err := time.Parse("2006-01-02", request.ReservationDate)

	transaction := domain.Transaction{
		ID:                  uuid.New(),
		UserID:              request.UserID,
		TouristAttractionID: request.TouristAttractionID,
		Qty:                 request.Qty,
		Amount:              touristAttraction.TicketPrice * float64(request.Qty),
		ReservationDate:     reservationDate,
		Status:              "PENDING",
	}
	// Convert string to time.Time
	transaction.ReservationDate, err = time.Parse("2006-01-02", request.ReservationDate)

	response, err := service.transactionRepository.Create(&transaction)
	if err != nil {
		return nil, err
	}

	payment := web.PaymentRequest{
		Amount:        transaction.Amount,
		TransactionID: response.ID,
	}

	result, err := service.paymentService.InitializePayment(&payment)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *TransactionServiceImpl) Delete(transactionID int) error {
	if err := service.transactionRepository.Delete(transactionID); err != nil {
		return err
	}
	return nil
}
