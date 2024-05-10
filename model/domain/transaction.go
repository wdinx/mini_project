package domain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID                  uuid.UUID         `gorm:"column:id;type:varchar(255);primaryKey;not null;default:uuid_generate_v4()"`
	UserID              uint              `gorm:"column:user_id"`
	User                User              `gorm:"foreignKey:UserID;references:ID"`
	TouristAttractionID int               `gorm:"column:tourist_attraction_id;not null"`
	TouristAttraction   TouristAttraction `gorm:"foreignKey:TouristAttractionID;references:ID"`
	Qty                 int               `gorm:"column:qty;not null"`
	Amount              float64           `gorm:"column:amount;not null"`
	ReservationDate     time.Time         `gorm:"column:reservation_date;not null;type:date"`
	Status              string            `gorm:"column:status;not null"`
}
