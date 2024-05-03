package domain

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name string `gorm:"column:name"`
	//TouristAttraction   TouristAttraction `gorm:"foreignKey:TouristAttractionID"`
	//TouristAttractionID int               `gorm:"column:tourist_attraction_id"`
	Username string `gorm:"column:username;unique;not null"`
	Password string `gorm:"column:password"`
}
