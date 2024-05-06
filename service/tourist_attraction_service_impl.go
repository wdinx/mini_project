package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util"
)

type TouristAttractionServiceImpl struct {
	touristAttractionRepository repository.TouristAttractionRepository
	validator                   *validator.Validate
}

func NewTouristAttractionService(touristAttractionRepository repository.TouristAttractionRepository, validator *validator.Validate) TouristAttractionService {
	return &TouristAttractionServiceImpl{touristAttractionRepository: touristAttractionRepository, validator: validator}
}

func (service TouristAttractionServiceImpl) Create(request *web.TouristAttractionRequest) (*web.TouristAttractionResponse, error) {
	if err := service.validator.Struct(request); err != nil {
		return nil, err
	}

	image := util.StoreImageToLocal(request.Image, request.Name)
	touristAttraction := domain.TouristAttraction{
		Name:        request.Name,
		TicketPrice: request.TicketPrice,
		Description: request.Description,
		Location:    request.Location,
		Image:       image,
		TypeID:      request.TouristAttractionTypeId,
	}

	result, err := service.touristAttractionRepository.Create(&touristAttraction)
	if err != nil {
		return nil, err
	}

	touristAttractionResponse := web.TouristAttractionResponse{
		Id:                    int(result.ID),
		Name:                  result.Name,
		TicketPrice:           result.TicketPrice,
		Description:           result.Description,
		Location:              result.Location,
		Image:                 result.Image,
		TouristAttractionType: result.TouristAttractionType,
		Balance:               result.Balance,
	}

	return &touristAttractionResponse, nil
}

func (service TouristAttractionServiceImpl) Update(request *web.TouristAttractionUpdateRequest) (*web.TouristAttractionResponse, error) {

	var image string
	if request.Image != nil {
		image = util.StoreImageToLocal(request.Image, request.Name)
	}
	touristAttraction := domain.TouristAttraction{
		ID:          request.ID,
		Name:        request.Name,
		TicketPrice: request.TicketPrice,
		Location:    request.Location,
		Image:       image,
		TypeID:      request.TouristAttractionTypeId,
	}

	result, err := service.touristAttractionRepository.Update(&touristAttraction)
	if err != nil {
		return nil, err
	}

	touristAttractionResponse := web.TouristAttractionResponse{
		Id:                    result.ID,
		Name:                  result.Name,
		TicketPrice:           result.TicketPrice,
		Location:              result.Location,
		Image:                 result.Image,
		TouristAttractionType: result.TouristAttractionType,
		Balance:               result.Balance,
	}
	return &touristAttractionResponse, nil
}

func (service TouristAttractionServiceImpl) UpdateBalanceById(request *web.TouristAttractionUpdateRequest) (*web.TouristAttractionBalanceResponse, error) {

	touristAttraction := domain.TouristAttraction{
		ID:      request.ID,
		Balance: request.Balance,
	}

	result, err := service.touristAttractionRepository.UpdateBalanceById(&touristAttraction)
	if err != nil {
		return nil, err
	}

	touristAttractionBalanceResponse := web.TouristAttractionBalanceResponse{
		Id:      result.ID,
		Balance: result.Balance,
	}
	return &touristAttractionBalanceResponse, nil
}

func (service TouristAttractionServiceImpl) GetAllTouristAttraction() (*[]web.TouristAttractionResponse, error) {

	result, err := service.touristAttractionRepository.GetAllTouristAttraction()
	if err != nil {
		return nil, err
	}

	var touristAttractionResponses []web.TouristAttractionResponse
	for _, touristAttraction := range *result {
		touristAttractionResponses = append(touristAttractionResponses, web.TouristAttractionResponse{
			Id:                    touristAttraction.ID,
			Name:                  touristAttraction.Name,
			TicketPrice:           touristAttraction.TicketPrice,
			Location:              touristAttraction.Location,
			Image:                 touristAttraction.Image,
			TouristAttractionType: touristAttraction.TouristAttractionType,
			Balance:               touristAttraction.Balance,
		})
	}
	return &touristAttractionResponses, nil
}
