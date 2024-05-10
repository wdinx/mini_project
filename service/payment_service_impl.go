package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"mini_project/model/domain"
	"mini_project/model/web"
	_interface2 "mini_project/repository/interface"
	"mini_project/service/interface"
)

type PaymentServiceImpl struct {
	paymentRepository        _interface2.PaymentRepository
	midtransService          _interface.MidtransService
	touristAttractionService _interface.TouristAttractionService
	ticketRepository         _interface2.TicketRepository
	transactionRepository    _interface2.TransactionRepository
}

func NewPaymentService(
	paymentRepository _interface2.PaymentRepository,
	midtransService _interface.MidtransService,
	touristAttractionService _interface.TouristAttractionService,
	ticketRepository _interface2.TicketRepository,
	transactionRepository _interface2.TransactionRepository) _interface.PaymentService {

	if paymentRepository == nil || midtransService == nil || touristAttractionService == nil || ticketRepository == nil || transactionRepository == nil {
		fmt.Println("PaymentService initialization failed")
	}

	return &PaymentServiceImpl{
		paymentRepository:        paymentRepository,
		midtransService:          midtransService,
		touristAttractionService: touristAttractionService,
		ticketRepository:         ticketRepository,
		transactionRepository:    transactionRepository,
	}
}

func (service *PaymentServiceImpl) InitializePayment(request *web.PaymentRequest) (*web.PaymentResponse, error) {
	payment := domain.Payment{
		ID:            uuid.New(),
		TransactionID: request.TransactionID,
		Amount:        request.Amount,
		Status:        0,
	}

	err := service.midtransService.GenerateSnapURL(&payment)
	if err != nil {
		return nil, err
	}

	if err = service.paymentRepository.Insert(&payment); err != nil {
		return nil, err
	}

	return &web.PaymentResponse{
		ID:      payment.ID,
		Amount:  payment.Amount,
		SnapURL: payment.SnapURL,
	}, nil
}

func (service *PaymentServiceImpl) ConfirmedPayment(id string) error {
	payment, err := service.paymentRepository.FindByID(id)
	if err != nil {
		return err
	}
	// Check if payment is found
	if *payment == (domain.Payment{}) {
		return errors.New("payment not found")
	}
	// Check if payment is already confirmed
	if payment.Status == 1 {
		return errors.New("payment already confirmed")
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
