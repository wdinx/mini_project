package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ID                  uuid.UUID         `gorm:"column:id;primaryKey;autoIncrement;not null"`
	TouristAttractionID int               `gorm:"column:tourist_attraction_id;not null"`
	TouristAttraction   TouristAttraction `gorm:"foreignKey:tourist_attraction_id;references:id"`
	UserID              uint              `gorm:"column:user_id;not null"`
	User                User              `gorm:"foreignKey:user_id;references:id"`
	TransactionID       uuid.UUID         `gorm:"column:transaction_id;not null"`
	Transaction         Transaction       `gorm:"foreignKey:transaction_id;references:id"`
	ReservationDate     time.Time         `gorm:"column:reservation_date;not null"`
}
