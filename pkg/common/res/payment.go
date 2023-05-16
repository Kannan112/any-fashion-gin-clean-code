package res

type RazorPayResponse struct {
	Email       string
	PhoneNumber string
	PaymentId   uint
	RazorpayKey string
	OrderId     interface{}
	AmountToPay uint
	Total       uint
}
