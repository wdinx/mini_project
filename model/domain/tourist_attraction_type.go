package domain

import "gorm.io/gorm"

type TouristAttractionType struct {
	gorm.Model
	Name string `gorm:"column:name;not null"`
}
