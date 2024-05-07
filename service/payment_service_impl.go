package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
)

type PaymentServiceImpl struct {
	paymentRepository           repository.PaymentRepository
	midtransService             MidtransService
	touristAttractionRepository repository.TouristAttractionRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository, midtransService MidtransService, touristAttractionRepository repository.TouristAttractionRepository) PaymentService {
	return &PaymentServiceImpl{
		paymentRepository:           paymentRepository,
		midtransService:             midtransService,
		touristAttractionRepository: touristAttractionRepository,
	}
}

func (service *PaymentServiceImpl) InitializePayment(request *web.PaymentRequest) (*web.PaymentResponse, error) {
	payment := domain.Payment{
		ID:                  uuid.NewString(),
		UserID:              request.UserID,
		TouristAttractionID: request.TouristAttractionID,
		Amount:              request.Amount,
		Status:              0,
	}
	err := service.midtransService.GenerateSnapURL(&payment)
	if err != nil {
		return nil, err
	}

	if err = service.paymentRepository.Insert(&payment); err != nil {
		return nil, err
	}

	userResponse := web.UserLoginResponse{
		ID:    payment.UserID,
		Name:  payment.User.Name,
		Email: payment.User.Email,
	}

	touristAttraction := web.TouristAttractionResponse{
		Id:          payment.TouristAttractionID,
		Name:        payment.TouristAttraction.Name,
		Description: payment.TouristAttraction.Description,
		Location:    payment.TouristAttraction.Location,
		TicketPrice: payment.TouristAttraction.TicketPrice,
		Image:       payment.TouristAttraction.Image,
	}

	return &web.PaymentResponse{
		ID:                payment.ID,
		Amount:            payment.Amount,
		User:              userResponse,
		TouristAttraction: touristAttraction,
		SnapURL:           payment.SnapURL,
	}, nil
}

func (service *PaymentServiceImpl) ConfirmedPayment(id string) error {
	fmt.Println(id)
	payment, err := service.paymentRepository.FindByID(id)
	if err != nil {
		fmt.Println("data payment tidak ditemukan")
		return err
	}

	if *payment == (domain.Payment{}) {
		fmt.Println("data payment tidak ada")
		return errors.New("payment not found")
	}

	fmt.Println(payment.TouristAttractionID)

	touristAttraction, err := service.touristAttractionRepository.FindByID(payment.TouristAttractionID)
	if err != nil {
		fmt.Println("data tempat wisata tidak ditemukan")
		return err
	}
	if *touristAttraction == (domain.TouristAttraction{}) {
		fmt.Println("data tempat wisata tidak ada")
		return errors.New("tourist attraction not found")
	}
	fmt.Println(touristAttraction.Balance, "and", payment.Amount)
	touristAttraction.Balance += int(payment.Amount)
	fmt.Println(touristAttraction.Balance)
	data, err := service.touristAttractionRepository.Update(touristAttraction)
	if err != nil {
		fmt.Println("data payment tidak berhasil di update", data)
		return err
	}
	fmt.Println("data berhasil di update")
	return nil
}
