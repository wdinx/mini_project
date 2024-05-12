package domain

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalRecord int `json:"total_record"`
}
