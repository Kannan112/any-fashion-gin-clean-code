package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	"golang.org/x/net/context"
)

type CouponRepository interface {
	FindCouponByName(ctx context.Context, couponCode string) (bool, error)
	AddCoupon(ctx context.Context, coupon req.Coupons) error
	UpdateCoupon(ctx context.Context, coupon req.Coupons, CouponId int) error
	DeleteCoupon(ctx context.Context, Couponid int) error
	ViewCoupon(ctx context.Context) ([]domain.Coupon, error)
	ApplyCoupon(ctx context.Context, userId int, couponCode string) (int, error)
}
