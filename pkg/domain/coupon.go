package domain

import "time"

type Coupon struct {
	Id                    uint      `gorm:"primeryKey;not null"`
	Code                  string    `json:"code"`
	DiscountPercent       float64   `json:"discountpercent"`
	DiscountMaximumAmount float64   `json:"discountmaximumamount"`
	MinimumPurchaseAmount float64   `json:"minimumpurchaseamount"`
	ExpirationDate        time.Time `json:"expirationdate"`
}
