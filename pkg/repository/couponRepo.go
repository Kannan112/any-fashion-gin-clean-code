package repository

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type CouponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &CouponDatabase{DB}
}

func (c *CouponDatabase) AddCoupon(ctx context.Context, coupon req.Coupon) error {
	query := `INSERT INTO coupons (code,discount_percentage,maximum_discount,minimum_purchase_amount,expire)VALUES($1,$2,$3,$4,$5)`
	err := c.DB.Exec(query, coupon.Code, coupon.DiscountPercentage, coupon.MaximumDiscount, coupon.MinimumPurchaseAmount, coupon.Expire).Error
	return err
}
