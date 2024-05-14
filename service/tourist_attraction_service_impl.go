package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/constant"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util"
	"mini_project/util/converter"
)

type TouristAttractionServiceImpl struct {
	touristAttractionRepository repository.TouristAttractionRepository
	imageService                ImageService
	validator                   *validator.Validate
}

func NewTouristAttractionService(touristAttractionRepository repository.TouristAttractionRepository, imageService ImageService, validator *validator.Validate) TouristAttractionService {
	return &TouristAttractionServiceImpl{touristAttractionRepository: touristAttractionRepository, imageService: imageService, validator: validator}
}

func (service *TouristAttractionServiceImpl) Create(request *web.TouristAttractionRequest) (*web.TouristAttractionResponse, error) {
	if err := service.validator.Struct(request); err != nil {
		return &web.TouristAttractionResponse{}, constant.ErrEmptyInput
	}

	filename := util.GenerateImageName(request.Name, request.Image.Filename)

	err := service.imageService.UploadImage(request.Image, filename)
	if err != nil {
		return &web.TouristAttractionResponse{}, err
	}

	touristAttraction := converter.ToTouristAttractionModel(request, filename)

	result, err := service.touristAttractionRepository.Create(touristAttraction)
	if err != nil {
		return &web.TouristAttractionResponse{}, err
	}

	touristAttractionResponse := converter.ToTouristAttractionResponse(result)

	return touristAttractionResponse, nil
}

func (service *TouristAttractionServiceImpl) Update(request *web.TouristAttractionUpdateRequest) (*web.TouristAttractionResponse, error) {

	filename := util.GenerateImageName(request.Name, request.Image.Filename)

	err := service.imageService.UploadImage(request.Image, filename)
	if err != nil {
		return &web.TouristAttractionResponse{}, err
	}

	touristAttraction := converter.ToUpdateTouristAttractionModel(request, filename)

	result, err := service.touristAttractionRepository.Update(touristAttraction)
	if err != nil {
		return &web.TouristAttractionResponse{}, err
	}

	touristAttractionResponse := converter.ToTouristAttractionResponse(result)
	return touristAttractionResponse, nil
}

func (service *TouristAttractionServiceImpl) UpdateBalanceById(request *web.TouristAttractionUpdateRequest) error {

	touristAttraction, err := service.touristAttractionRepository.FindByID(request.ID)
	if err != nil {
		return err
	}

	touristAttraction.Balance += request.Balance

	_, err = service.touristAttractionRepository.UpdateBalanceById(touristAttraction)
	if err != nil {
		return err
	}

	return nil
}

func (service *TouristAttractionServiceImpl) GetAllTouristAttraction() (*[]web.TouristAttractionResponse, error) {

	result, err := service.touristAttractionRepository.GetAllTouristAttraction()
	if err != nil {
		return &[]web.TouristAttractionResponse{}, err
	}

	var touristAttractionResponses []web.TouristAttractionResponse
	for _, touristAttraction := range *result {
		touristAttractionResponses = append(touristAttractionResponses, *converter.ToTouristAttractionResponse(&touristAttraction))
	}
	return &touristAttractionResponses, nil
}

func (service *TouristAttractionServiceImpl) GetTouristAttractionById(id int) (*web.TouristAttractionResponse, error) {
	result, err := service.touristAttractionRepository.FindByID(id)
	if err != nil {
		return &web.TouristAttractionResponse{}, err
	}

	touristAttractionResponse := converter.ToTouristAttractionResponse(result)
	return touristAttractionResponse, nil
}
