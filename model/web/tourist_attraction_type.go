package web

type TouristAttractionTypeCreateRequest struct {
	Name string `json:"name" validate:"required,min=1,max=200"`
}

type TouristAttractionTypeUpdateRequest struct {
	ID   uint
	Name string `json:"name" validate:"required,min=1,max=200"`
}

type TouristAttractionTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
