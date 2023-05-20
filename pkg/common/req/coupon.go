package req

import (
	"time"
)

type Coupon struct {
	Code                  string
	DiscountPercentage    float64
	MaximumDiscount       float64
	MinimumPurchaseAmount float64
	Expire                time.Time
}
