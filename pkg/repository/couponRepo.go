package repository

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type CouponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &CouponDatabase{DB}
}

func (c *CouponDatabase) AddCoupon(ctx context.Context, coupon req.Coupons) error {
	query := `INSERT INTO coupons (code,discount_percent,discount_maximum_amount,minimum_purchase_amount, expiration_date)VALUES($1,$2,$3,$4,$5)`
	err := c.DB.Exec(query, coupon.Code, coupon.DiscountPercent, coupon.DiscountMaximumAmount, coupon.MinimumPurchaseAmount, coupon.ExpirationDate).Error
	return err
}
func (c *CouponDatabase) UpdateCoupon(ctx context.Context, coupon req.Coupons, CouponId int) error {

	query := `UPDATE coupons 
	SET code=$1, 
		discount_percent=$2,
		discount_maximum_amount=$3,
		minimum_purchase_amount=$4,
		expiration_date=$5 
	WHERE id=$6;
	`
	err := c.DB.Exec(query, coupon.Code, coupon.DiscountPercent, coupon.DiscountMaximumAmount, coupon.MinimumPurchaseAmount, coupon.ExpirationDate, CouponId).Error
	return err
}
func (c *CouponDatabase) DeleteCoupon(ctx context.Context, couponId int) error {
	fmt.Println("couponId,", couponId)
	var check bool
	Querycheck := `SELECT EXISTS(SELECT 1 FROM coupons WHERE id=$1)`
	err := c.DB.Raw(Querycheck, couponId).Scan(&check).Error
	if err != nil {
		return err
	}
	if !check {
		return fmt.Errorf("Coupon with id not exists")
	}
	query := `DELETE FROM coupons WHERE id=$1`
	err = c.DB.Exec(query, couponId).Error
	return err
}
func (c *CouponDatabase) ViewCoupon(ctx context.Context) ([]domain.Coupon, error) {
	var coupon []domain.Coupon
	query := `SELECT * FROM coupons`
	err := c.DB.Raw(query).Scan(&coupon).Error
	return coupon, err
}
