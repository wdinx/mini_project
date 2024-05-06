package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
	"strconv"
)

type TransactionServiceImpl struct {
	transactionRepository       repository.TransactionRepository
	touristAttractionRepository repository.TouristAttractionRepository
	validator                   *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, touristAttractionRepository repository.TouristAttractionRepository, validator *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		transactionRepository:       transactionRepository,
		touristAttractionRepository: touristAttractionRepository,
		validator:                   validator,
	}
}

func (service TransactionServiceImpl) Create(request *web.TransactionCreateRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}

	touristAttraction, err := service.touristAttractionRepository.FindByID(request.TouristAttractionID)
	if err != nil {
		return err
	}

	ticketPrice, err := strconv.Atoi(touristAttraction.TicketPrice)
	if err != nil {
		return err
	}

	transaction := domain.Transaction{
		UserID:              request.UserID,
		TouristAttractionID: request.TouristAttractionID,
		Qty:                 request.Qty,
		Amount:              float64(ticketPrice * request.Qty),
	}

	if err := service.transactionRepository.Create(&transaction); err != nil {
		return err
	}
	return err
}

func (service TransactionServiceImpl) Delete(transactionID int) error {
	if err := service.transactionRepository.Delete(transactionID); err != nil {
		return err
	}
	return nil
}
