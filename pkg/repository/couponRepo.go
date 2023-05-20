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
func (c *CouponDatabase) UpdateCoupon(ctx context.Context, coupon req.Coupon, CouponId int) error {
	query := `UPDATE coupons SET (code=$1,discount_percentage=$2,maximum_discount=$3,minimum_purchase_amount=$4,expire=$5) WHERE id=$6`
	err := c.DB.Exec(query, coupon.Code, coupon.DiscountPercentage, coupon.MaximumDiscount, coupon.MinimumPurchaseAmount, coupon.Expire, CouponId).Error
	return err

}

func (c *CouponDatabase) DeleteCoupon(ctx context.Context, couponId int) error {
	query := `DELETE FROM coupons WHERE id $1`
	err := c.DB.Exec(query, couponId).Error
	return err
}
