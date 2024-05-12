package converter

import (
	"mini_project/model/domain"
	"mini_project/model/web"
	"mini_project/util"
)

func ToTouristAttractionModel(request *web.TouristAttractionRequest) *domain.TouristAttraction {
	image := util.StoreImageToLocal(request.Image, request.Name)
	return &domain.TouristAttraction{
		Name:        request.Name,
		Location:    request.Location,
		Description: request.Description,
		TypeID:      request.TouristAttractionTypeId,
		TicketPrice: request.TicketPrice,
		Image:       image,
	}
}

func ToUpdateTouristAttractionModel(request *web.TouristAttractionUpdateRequest) *domain.TouristAttraction {
	var image string
	if request.Image != nil {
		image = util.StoreImageToLocal(request.Image, request.Name)
	}
	return &domain.TouristAttraction{
		ID:          request.ID,
		Name:        request.Name,
		TypeID:      request.TouristAttractionTypeId,
		Description: request.Description,
		TicketPrice: request.TicketPrice,
		Location:    request.Location,
		Image:       image,
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
		Image:                 touristAttraction.Image,
		Balance:               touristAttraction.Balance,
	}
}
