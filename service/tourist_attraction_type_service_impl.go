package service

import (
	"github.com/go-playground/validator/v10"
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/repository"
)

type TouristAttractionTypeServiceImpl struct {
	touristAttractionTypeRepository repository.TouristAttractionTypeRepository
	validate                        *validator.Validate
}

func NewTouristAttractionTypeService(touristAttractionTypeRepository repository.TouristAttractionTypeRepository, validate *validator.Validate) TouristAttractionTypeService {
	return &TouristAttractionTypeServiceImpl{touristAttractionTypeRepository: touristAttractionTypeRepository, validate: validate}
}

func (service TouristAttractionTypeServiceImpl) Create(request *web.TouristAttractionTypeCreateRequest) (*web.TouristAttractionTypeResponse, error) {
	var err error
	if err = service.validate.Struct(request); err != nil {
		return nil, err
	}

	touristAttractionType := domain.TouristAttractionType{
		Name: request.Name,
	}

	result, err := service.touristAttractionTypeRepository.Create(&touristAttractionType)
	if err != nil {
		return nil, err
	}

	response := web.TouristAttractionTypeResponse{
		ID:   result.ID,
		Name: result.Name,
	}

	return &response, nil

}

func (service TouristAttractionTypeServiceImpl) Update(request *web.TouristAttractionTypeUpdateRequest) (*web.TouristAttractionTypeResponse, error) {
	var err error
	if err = service.validate.Struct(request); err != nil {
		return nil, err
	}

	data, err := service.touristAttractionTypeRepository.FindByID(int(request.ID))
	if err != nil {
		return nil, err
	}

	data.Name = request.Name

	result, err := service.touristAttractionTypeRepository.Update(data)
	if err != nil {
		return nil, err
	}

	response := web.TouristAttractionTypeResponse{
		ID:   result.ID,
		Name: result.Name,
	}

	return &response, nil
}

func (service TouristAttractionTypeServiceImpl) Delete(touristAttractionTypeId int) error {
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

func (service TouristAttractionTypeServiceImpl) GetAll() *[]web.TouristAttractionTypeResponse {
	touristAttractionTypes := service.touristAttractionTypeRepository.GetAll()
	var responses []web.TouristAttractionTypeResponse

	for _, touristAttractionType := range touristAttractionTypes {
		response := web.TouristAttractionTypeResponse{
			ID:   touristAttractionType.ID,
			Name: touristAttractionType.Name,
		}
		responses = append(responses, response)
	}

	return &responses
}