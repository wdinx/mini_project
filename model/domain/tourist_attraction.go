package domain

import "gorm.io/gorm"

type TouristAttraction struct {
	gorm.Model
	TypeID                int                   `gorm:"column:type_id;not null"`
	TouristAttractionType TouristAttractionType `gorm:"foreignKey:TypeID"`
	Name                  string                `gorm:"column:name;not null"`
	TicketPrice           string                `gorm:"column:ticket_price;not null"`
	Image                 string                `gorm:"column:image;not null"`
}
