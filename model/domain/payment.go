package domain

type Payment struct {
	ID                  string            `gorm:"column:id;primaryKey;autoIncrement;not null"`
	UserID              int               `gorm:"column:user_id;not null"`
	User                User              `gorm:"foreignKey:user_id;references:id"`
	TouristAttractionID int               `gorm:"column:tourist_attraction_id;not null"`
	TouristAttraction   TouristAttraction `gorm:"foreignKey:TouristAttractionID;references:id"`
	Amount              float64           `gorm:"column:amount;not null"`
	Status              int               `gorm:"column:status;not null"`
	SnapURL             string            `gorm:"column:snap_url;not null"`
}
