package repository

import (
	"context"
	"fmt"
	"time"

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

func (c *CouponDatabase) FindCouponByName(ctx context.Context, couponCode string) (bool, error) {
	var check bool
	query := `SELECT EXISTS(SELECT code FROM coupons WHERE code=$1)`
	err := c.DB.Raw(query, couponCode).Scan(&check).Error
	if err != nil {
		return false, err
	}
	return check, nil
}

func (c *CouponDatabase) AddCoupon(ctx context.Context, coupon req.Coupons) error {
	query := `INSERT INTO coupons (code, discount_percent, discount_maximum_amount, minimum_purchase_amount, expiration_date)
		VALUES ($1, $2, $3, $4, $5)`
	err := c.DB.Exec(query, coupon.Code, coupon.DiscountPercent, coupon.DiscountMaximumAmount, coupon.MinimumPurchaseAmount, coupon.ExpirationDate).Error
	return err
}

func (c *CouponDatabase) UpdateCoupon(ctx context.Context, coupon req.UpdateCoupon, couponID int) error {
	var check bool
	tx := c.DB.Begin()
	query := `SELECT EXISTS(SELECT code FROM coupons WHERE id=$1)`
	err := c.DB.Raw(query, couponID).Scan(&check).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if !check {
		tx.Rollback()
		return fmt.Errorf("no coupon code found")
	}

	updateQuery := `UPDATE coupons
		SET discount_percent=$1,
			discount_maximum_amount=$2,
			minimum_purchase_amount=$3,
			expiration_date=$4
		WHERE id=$5`
	err = c.DB.Exec(updateQuery, coupon.DiscountPercent, coupon.DiscountMaximumAmount, coupon.MinimumPurchaseAmount, coupon.ExpirationDate, couponID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (c *CouponDatabase) DeleteCoupon(ctx context.Context, couponID int) error {
	query := `DELETE FROM coupons WHERE id=$1`
	err := c.DB.Exec(query, couponID).Error
	return err
}

func (c *CouponDatabase) ViewCoupon(ctx context.Context) ([]domain.Coupon, error) {
	var coupons []domain.Coupon
	query := `SELECT * FROM coupons`
	err := c.DB.Raw(query).Scan(&coupons).Error
	return coupons, err
}

func (c *CouponDatabase) ApplyCoupon(ctx context.Context, userID int, couponCode string) (int, error) {
	tx := c.DB.Begin()
	var check bool
	// Check if coupon exists
	couponExistsQuery := `SELECT EXISTS(SELECT * FROM coupons WHERE code=$1)`
	err := tx.Raw(couponExistsQuery, couponCode).Scan(&check).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if !check {
		tx.Rollback()
		return 0, fmt.Errorf("coupon not found")
	}

	// Get the coupon details
	var coupon domain.Coupon
	couponQuery := `SELECT * FROM coupons WHERE code=$1`
	err = tx.Raw(couponQuery, couponCode).Scan(&coupon).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Check if coupon has expired
	if coupon.ExpirationDate.Before(time.Now()) {
		tx.Rollback()
		return 0, fmt.Errorf("coupon expired")
	}

	// Check if the coupon is already used
	var couponUsed bool
	couponUsedQuery := `SELECT EXISTS(SELECT 1 FROM orders WHERE coupon_code=$1)`
	err = tx.Raw(couponUsedQuery, couponCode).Scan(&couponUsed).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if couponUsed {
		tx.Rollback()
		return 0, fmt.Errorf("coupon has already been used")
	}

	// Check if the coupon is already applied to the cart
	var cartDetails domain.Carts
	getCartDetailsQuery := `SELECT * FROM carts WHERE users_id=$1`
	err = tx.Raw(getCartDetailsQuery, userID).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if cartDetails.CouponId == coupon.Id {
		tx.Rollback()
		return 0, fmt.Errorf("coupon is already applied to the cart")
	}

	// Check if the cart is empty
	if cartDetails.Sub_total == 0 {
		tx.Rollback()
		return 0, fmt.Errorf("cart is empty")
	}

	// Check if the coupon can be applied for the minimum purchase amount
	if cartDetails.Sub_total <= int(coupon.MinimumPurchaseAmount) {
		tx.Rollback()
		return 0, fmt.Errorf("minimum purchase amount should be %v", coupon.MinimumPurchaseAmount)
	}

	// Calculate the discount amount
	discountAmount := (cartDetails.Sub_total / 100) * int(coupon.DiscountPercent)
	if discountAmount > int(coupon.DiscountMaximumAmount) {
		discountAmount = int(coupon.DiscountMaximumAmount)
	}

	// Update the cart total by subtracting the discount amount
	updateCartQuery := `UPDATE carts SET total=$1, coupon_id=$2 WHERE id=$3`
	err = tx.Exec(updateCartQuery, cartDetails.Sub_total-discountAmount, coupon.Id, cartDetails.Id).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return cartDetails.Total - cartDetails.Sub_total, nil
}

func (c *CouponDatabase) RemoveCoupon(ctx context.Context, userID int) error {
	tx := c.DB.Begin()
	var cartDetails domain.Carts
	var couponDetails domain.Coupon
	getCartDetailsQuery := `SELECT * FROM carts WHERE users_id=$1`
	err := tx.Raw(getCartDetailsQuery, userID).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if cartDetails.CouponId == 0 {
		return fmt.Errorf("coupon is not used")
	}

	couponDetailsQuery := `SELECT * FROM coupons WHERE id=$1`
	err = tx.Raw(couponDetailsQuery, cartDetails.CouponId).Scan(&couponDetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if couponDetails.ExpirationDate.Before(time.Now()) {
		return fmt.Errorf("coupon has expired")
	}

	discountAmount := (cartDetails.Sub_total / 100) * int(couponDetails.DiscountPercent)
	if discountAmount > int(couponDetails.DiscountMaximumAmount) {
		discountAmount = int(couponDetails.DiscountMaximumAmount)
	}

	updateCartQuery := `UPDATE carts SET coupon_id=$1, total=$2 WHERE id=$3`
	err = tx.Exec(updateCartQuery, nil, discountAmount+cartDetails.Total, cartDetails.Id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
