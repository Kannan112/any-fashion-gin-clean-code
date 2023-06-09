package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type CouponUseCase struct {
	coupoRepo interfaces.CouponRepository
}

func NewCouponUseCase(couponRepo interfaces.CouponRepository) services.CouponUseCase {
	return &CouponUseCase{couponRepo}
}
func (c *CouponUseCase) AddCoupon(ctx context.Context, coupon req.Coupons) error {
	checkCoupon, err := c.coupoRepo.FindCouponByName(ctx, coupon.Code)
	if err != nil {
		return err
	} else if checkCoupon {
		return fmt.Errorf("there already a coupon exist with coupon_name %v", coupon.Code)
	}
	// validate the coupn expire date
	if time.Since(coupon.ExpirationDate) > 0 {
		return fmt.Errorf("given coupon expire date already exceeded %v", coupon.ExpirationDate)
	}
	err = c.coupoRepo.AddCoupon(ctx, coupon)
	return err
}
func (c *CouponUseCase) DeleteCoupon(ctx context.Context, couponId int) error {
	err := c.coupoRepo.DeleteCoupon(ctx, couponId)
	return err
}
func (c *CouponUseCase) UpdateCoupon(ctx context.Context, coupon req.UpdateCoupon, CouponId int) error {
	err := c.coupoRepo.UpdateCoupon(ctx, coupon, CouponId)
	return err
}
func (c *CouponUseCase) ViewCoupon(ctx context.Context) ([]domain.Coupon, error) {
	coupon, err := c.coupoRepo.ViewCoupon(ctx)
	return coupon, err
}
func (c *CouponUseCase) ApplyCoupon(ctx context.Context, userId int, couponCode string) (int, error) {
	total, err := c.coupoRepo.ApplyCoupon(ctx, userId, couponCode)
	return total, err
}
func (c *CouponUseCase) RemoveCoupon(ctx context.Context, userId int) error {
	err := c.coupoRepo.RemoveCoupon(ctx, userId)
	return err
}

