package domain

type Transaction struct {
	ID                  string            `json:"id" gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4();not null" validate:"required"`
	UserID              int               `json:"user_id" gorm:"column:user_id;not null" validate:"required"`
	User                User              `gorm:"foreignKey:UserID;references:ID"`
	TouristAttractionID int               `json:"tourist_attraction_id" gorm:"column:tourist_attraction_id;not null" validate:"required"`
	TouristAttraction   TouristAttraction `gorm:"foreignKey:TouristAttractionID;references:ID"`
	Qty                 int               `json:"qty" gorm:"column:qty;not null" validate:"required"`
	Amount              float64           `json:"amount" gorm:"column:amount;not null"`
}
