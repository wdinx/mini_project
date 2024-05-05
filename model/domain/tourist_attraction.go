package domain

type TouristAttraction struct {
	ID                    int                   `gorm:"column:id;primaryKey;autoIncrement"`
	TypeID                int                   `gorm:"column:type_id;not null"`
	TouristAttractionType TouristAttractionType `gorm:"foreignKey:type_id;references:id"`
	Description           string                `gorm:"column:description;not null"`
	Name                  string                `gorm:"column:name;not null"`
	TicketPrice           string                `gorm:"column:ticket_price;not null"`
	Location              string                `gorm:"column:location;not null"`
	Image                 string                `gorm:"column:image;not null"`
	Balance               int                   `gorm:"column:balance;not null;default:0"`
}
