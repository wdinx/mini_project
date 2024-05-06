package domain

type Payment struct {
	ID            string      `gorm:"column:id;primaryKey;autoIncrement;type:uuid;default:uuid_generate_v4();not null"`
	TransactionID string      `gorm:"column:transaction_id;not null"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	Amount        float64     `gorm:"column:amount;not null"`
	Status        string      `gorm:"column:status;not null"`
	SnapURL       string      `gorm:"column:snap_url;not null"`
}
