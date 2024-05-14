package service

import (
	"github.com/google/uuid"
	"mini_project/constant"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util/converter"
)

type PaymentServiceImpl struct {
	paymentRepository        repository.PaymentRepository
	midtransService          MidtransService
	touristAttractionService TouristAttractionService
	ticketRepository         repository.TicketRepository
	transactionRepository    repository.TransactionRepository
}

func NewPaymentService(
	paymentRepository repository.PaymentRepository,
	midtransService MidtransService,
	touristAttractionService TouristAttractionService,
	ticketRepository repository.TicketRepository,
	transactionRepository repository.TransactionRepository) PaymentService {

	return &PaymentServiceImpl{
		paymentRepository:        paymentRepository,
		midtransService:          midtransService,
		touristAttractionService: touristAttractionService,
		ticketRepository:         ticketRepository,
		transactionRepository:    transactionRepository,
	}
}

func (service *PaymentServiceImpl) InitializePayment(request *web.PaymentRequest) (*web.PaymentResponse, error) {
	payment := converter.ToPaymentModel(request)

	err := service.midtransService.GenerateSnapURL(payment)
	if err != nil {
		return &web.PaymentResponse{}, err
	}

	if err = service.paymentRepository.Insert(payment); err != nil {
		return nil, err
	}

	return converter.ToPaymentResponse(payment), nil
}

func (service *PaymentServiceImpl) ConfirmedPayment(id string) error {
	payment, err := service.paymentRepository.FindByID(id)
	if err != nil {
		return err
	}
	// Check if payment is found
	if *payment == (domain.Payment{}) {
		return constant.ErrPaymentNotFound
	}
	// Check if payment is already confirmed
	if payment.Status == 1 {
		return constant.ErrPaymentAlreadyConfirmed
	}

	// Action to update balance
	updateTouristAttraction := web.TouristAttractionUpdateRequest{
		ID:      payment.Transaction.TouristAttractionID,
		Balance: payment.Amount,
	}
	err = service.touristAttractionService.UpdateBalanceById(&updateTouristAttraction)
	if err != nil {
		return err
	}

	// Action to update payment status
	payment.Status = 1
	err = service.paymentRepository.Update(payment)
	if err != nil {
		return err
	}

	// Action to create ticket
	ticket := domain.Ticket{
		ID:                  uuid.New(),
		TouristAttractionID: payment.Transaction.TouristAttractionID,
		UserID:              payment.Transaction.UserID,
		TransactionID:       payment.Transaction.ID,
		ReservationDate:     payment.Transaction.ReservationDate,
	}
	if err = service.ticketRepository.Insert(&ticket); err != nil {
		return err
	}

	return nil
}
