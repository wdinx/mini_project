package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"mini_project/constant"
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
		return &web.PaymentResponse{}, constant.ErrEmptyInput
	}

	touristAttraction, err := service.touristAttractionRepository.FindByID(request.TouristAttractionID)
	if err != nil {
		return &web.PaymentResponse{}, err
	}

	reservationDate, err := util.StringToDate(request.ReservationDate)
	if err != nil {
		return &web.PaymentResponse{}, err
	}

	transaction := converter.ToTransactionModel(request, touristAttraction, reservationDate)

	response, err := service.transactionRepository.Create(transaction)
	if err != nil {
		return &web.PaymentResponse{}, err
	}

	payment := web.PaymentRequest{
		Amount:        transaction.Amount,
		TransactionID: response.ID,
	}

	result, err := service.paymentService.InitializePayment(&payment)
	if err != nil {
		return &web.PaymentResponse{}, err
	}
	return result, nil
}

func (service *TransactionServiceImpl) Delete(transactionID int) error {
	if err := service.transactionRepository.Delete(transactionID); err != nil {
		return err
	}
	return nil
}

func (service *TransactionServiceImpl) GetByUserID(userID int) (*[]web.TransactionResponse, error) {
	transactions, err := service.transactionRepository.GetByUserID(userID)
	if err != nil {
		return &[]web.TransactionResponse{}, err
	}

	if len(*transactions) == 0 {
		return &[]web.TransactionResponse{}, constant.ErrDataNotFound
	}

	var responses []web.TransactionResponse
	for _, transaction := range *transactions {
		response := converter.ToTransactionResponse(&transaction)
		fmt.Println(*response)
		responses = append(responses, *response)
	}
	return &responses, nil
}

func (service *TransactionServiceImpl) GetByID(transactionID int) (*web.TransactionResponse, error) {
	transaction, err := service.transactionRepository.GetByID(transactionID)
	if err != nil {
		return &web.TransactionResponse{}, err
	}

	response := converter.ToTransactionResponse(transaction)
	return response, nil
}
