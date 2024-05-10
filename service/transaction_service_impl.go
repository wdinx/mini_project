package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util"
	"mini_project/util/converter"
)

type TransactionServiceImpl struct {
	transactionRepository       repository.TransactionRepository
	touristAttractionRepository repository.TouristAttractionRepository
	paymentService              PaymentService
	validator                   *validator.Validate
}

func NewTransactionService(
	transactionRepository repository.TransactionRepository,
	touristAttractionRepository repository.TouristAttractionRepository,
	paymentService PaymentService,
	validator *validator.Validate) TransactionService {
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

	request.ReservationDate, err = util.StringToDate(request.ReservationDate.String())
	if err != nil {
		return nil, err
	}

	transaction := converter.ToTransactionModel(request, touristAttraction)

	response, err := service.transactionRepository.Create(transaction)
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
