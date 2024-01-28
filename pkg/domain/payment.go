package domain

type PaymentType string

const (
	// payment type
	RazopayPayment      PaymentType = "razor payment"
	RazorayMaxiumAmount             = 50000
	CodePayment         PaymentType = "cash on delivery"
	CodMaxiumAmount                 = 20000
	StripePayment       PaymentType = "stripe payment"
	StripeMaxiumAmount              = 35000
)

type PaymentMethod struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	Name          PaymentType `json:"name" gorm:"unique"`
	BlockStatus   bool        `json:"block_status" gorm:"default:false"`
	MaximumAmount uint        `json:"maximum_amount"`
}
