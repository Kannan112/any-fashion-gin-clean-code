package res

import "github.com/kannan112/go-gin-clean-arch/pkg/domain"

type RazorPayResponse struct {
	Email       string
	PhoneNumber string
	PaymentId   uint
	RazorpayKey string
	OrderId     interface{}
	AmountToPay uint
	Total       uint
}

type StripeOrder struct {
	ClientSecret   string `json:"client_secret"`
	PublishableKey string `json:"publishable_key"`
	AmountToPay    uint   `json:"amount_to_pay"`
	CartID         uint   `json:"cart_id"`
}

type OrderPayment struct {
	PaymentType  domain.PaymentType `json:"payment_type"`
	PaymentOrder any                `json:"payment_order"`
}

type RazorpayOrder struct {
	RazorpayKey     string      `json:"razorpay_key"`
	UserID          uint        `json:"user_id"`
	AmountToPay     uint        `json:"amount_to_pay"`
	RazorpayAmount  uint        `json:"razorpay_amount"`
	RazorpayOrderID interface{} `json:"razorpay_order_id"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`

	ShopOrderID uint `json:"shop_order_id"`
}
