package domain

import "github.com/google/uuid"

type Payment struct {
	ID            uuid.UUID   `gorm:"column:id;type:varchar(255);primaryKey;not null;"`
	TransactionID uuid.UUID   `gorm:"column:transaction_id;not null"`
	Transaction   Transaction `gorm:"foreignKey:transaction_id;references:id"`
	Amount        float64     `gorm:"column:amount;not null"`
	Status        int         `gorm:"column:status;not null;default:0"`
	SnapURL       string      `gorm:"column:snap_url;not null"`
}
