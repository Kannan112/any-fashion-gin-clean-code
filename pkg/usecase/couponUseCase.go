package usecase

import (
	"context"

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
	err := c.coupoRepo.AddCoupon(ctx, coupon)
	return err
}
func (c *CouponUseCase) DeleteCoupon(ctx context.Context, couponId int) error {
	err := c.coupoRepo.DeleteCoupon(ctx, couponId)
	return err
}
func (c *CouponUseCase) UpdateCoupon(ctx context.Context, coupon req.Coupons, CouponId int) error {
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
