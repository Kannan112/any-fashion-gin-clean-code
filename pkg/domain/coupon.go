package domain

import "time"

type Coupon struct {
	Id                    uint `gorm:"primeryKey;not null"`
	Code                  string
	DiscountPercentage    float64
	MaximumDiscount       float64
	MinimumPurchaseAmount float64
	Expire                time.Time
}
