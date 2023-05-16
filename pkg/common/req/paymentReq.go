package req

type PaymentReq struct {
	PaymentType   string `json:"payment_type" gorm:"unique;not null"`
	BlockStatus   bool   `json:"block_status" gorm:"not null;default:false"`
	MaximumAmount uint   `json:"maximum_amount" gorm:"not null"`
}

type RazorPayRequest struct {
	RazorPayPaymentId  string
	RazorPayOrderId    string
	Razorpay_signature string
}
