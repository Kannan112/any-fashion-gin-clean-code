package usecase

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type CouponUseCase struct {
	coupoRepo interfaces.CouponRepository
}

func NewCouponUseCase(couponRepo interfaces.CouponRepository) services.CouponUseCase {
	return &CouponUseCase{couponRepo}
}
func (c *CouponUseCase) AddCoupon(ctx context.Context, coupon req.Coupon) error {
	err := c.coupoRepo.AddCoupon(ctx, coupon)
	return err
}
