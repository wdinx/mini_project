package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/constant"
	"mini_project/model/web"
	"mini_project/repository"
	"mini_project/util/converter"
)

type TouristAttractionTypeServiceImpl struct {
	touristAttractionTypeRepository repository.TouristAttractionTypeRepository
	validator                       *validator.Validate
}

func NewTouristAttractionTypeService(touristAttractionTypeRepository repository.TouristAttractionTypeRepository, validate *validator.Validate) TouristAttractionTypeService {
	return &TouristAttractionTypeServiceImpl{touristAttractionTypeRepository: touristAttractionTypeRepository, validator: validate}
}

func (service *TouristAttractionTypeServiceImpl) Create(request *web.TouristAttractionTypeCreateRequest) (*web.TouristAttractionTypeResponse, error) {
	var err error
	if err = service.validator.Struct(request); err != nil {
		return &web.TouristAttractionTypeResponse{}, constant.ErrEmptyInput
	}

	touristAttractionType := converter.ToCreateTouristAttractionType(request)

	result, err := service.touristAttractionTypeRepository.Create(touristAttractionType)
	if err != nil {
		return &web.TouristAttractionTypeResponse{}, err
	}

	response := converter.ToTouristAttractionTypeResponse(result)

	return response, nil

}

func (service *TouristAttractionTypeServiceImpl) Update(request *web.TouristAttractionTypeUpdateRequest) (*web.TouristAttractionTypeResponse, error) {
	var err error
	if err = service.validator.Struct(request); err != nil {
		return &web.TouristAttractionTypeResponse{}, constant.ErrEmptyInput
	}

	data, err := service.touristAttractionTypeRepository.FindByID(int(request.ID))
	if err != nil {
		return &web.TouristAttractionTypeResponse{}, err
	}

	data.Name = request.Name

	result, err := service.touristAttractionTypeRepository.Update(data)
	if err != nil {
		return &web.TouristAttractionTypeResponse{}, err
	}

	response := converter.ToTouristAttractionTypeResponse(result)

	return response, nil
}

func (service *TouristAttractionTypeServiceImpl) Delete(touristAttractionTypeId int) error {
	_, err := service.touristAttractionTypeRepository.FindByID(touristAttractionTypeId)
	if err != nil {
		return err
	}

	err = service.touristAttractionTypeRepository.Delete(touristAttractionTypeId)
	if err != nil {
		return err
	}

	return nil
}

func (service *TouristAttractionTypeServiceImpl) GetAll() *[]web.TouristAttractionTypeResponse {
	touristAttractionTypes := service.touristAttractionTypeRepository.GetAll()
	var responses []web.TouristAttractionTypeResponse

	for _, touristAttractionType := range touristAttractionTypes {
		response := converter.ToTouristAttractionTypeResponse(&touristAttractionType)
		responses = append(responses, *response)
	}

	return &responses
}

func (service *TouristAttractionTypeServiceImpl) FindByID(touristAttractionTypeId int) (*web.TouristAttractionTypeResponse, error) {
	result, err := service.touristAttractionTypeRepository.FindByID(touristAttractionTypeId)
	if err != nil {
		return &web.TouristAttractionTypeResponse{}, err
	}

	response := converter.ToTouristAttractionTypeResponse(result)

	return response, nil
}
