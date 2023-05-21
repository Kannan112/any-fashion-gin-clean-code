package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type CouponUseCase interface {
	AddCoupon(ctx context.Context, coupon req.Coupons) error
	UpdateCoupon(ctx context.Context, coupon req.Coupons, CouponId int) error
	DeleteCoupon(ctx context.Context, couponId int) error
	ViewCoupon(ctx context.Context) ([]domain.Coupon, error)
}
