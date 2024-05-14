package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/util"
)

func ToTouristAttractionModel(request *web.TouristAttractionRequest, filename string) *domain.TouristAttraction {
	return &domain.TouristAttraction{
		Name:        request.Name,
		Location:    request.Location,
		Description: request.Description,
		TypeID:      request.TouristAttractionTypeId,
		TicketPrice: request.TicketPrice,
		Image:       filename,
	}
}

func ToUpdateTouristAttractionModel(request *web.TouristAttractionUpdateRequest, filename string) *domain.TouristAttraction {
	return &domain.TouristAttraction{
		ID:          request.ID,
		Name:        request.Name,
		TypeID:      request.TouristAttractionTypeId,
		Description: request.Description,
		TicketPrice: request.TicketPrice,
		Location:    request.Location,
		Image:       filename,
		Balance:     request.Balance,
	}
}

func ToTouristAttractionResponse(touristAttraction *domain.TouristAttraction) *web.TouristAttractionResponse {
	return &web.TouristAttractionResponse{
		Id:                    touristAttraction.ID,
		Name:                  touristAttraction.Name,
		Description:           touristAttraction.Description,
		TouristAttractionType: *ToTouristAttractionTypeResponse(&touristAttraction.TouristAttractionType),
		Location:              touristAttraction.Location,
		TicketPrice:           touristAttraction.TicketPrice,
		Image:                 util.GetImageUrl(touristAttraction.Image),
		Balance:               touristAttraction.Balance,
	}
}
