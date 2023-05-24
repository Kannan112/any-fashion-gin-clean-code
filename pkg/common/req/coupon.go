package req

import (
	"time"
)

type Coupons struct {
	Code                  string    `json:"code"`
	DiscountPercent       float64   `json:"discountpercent"`
	DiscountMaximumAmount float64   `json:"discountmaximumamount"`
	MinimumPurchaseAmount float64   `json:"minimumpurchaseamount"`
	ExpirationDate        time.Time `json:"expirationdate"`
}
	