package domain

type PaymentMethod struct {
	ID            uint   `json:"id" gorm:"primaryKey;not null"`
	PaymentType   string `json:"payment_type" gorm:"unique;not null"`
	BlockStatus   bool   `json:"block_status" gorm:"not null;default:false"`
	MaximumAmount uint   `json:"maximum_amount" gorm:"not null"`
	CreatedAt     string `json:"created_at" gorm:"not null"`
	UpdatedAt     string `json:"updated_at" gorm:"not null"`
}
