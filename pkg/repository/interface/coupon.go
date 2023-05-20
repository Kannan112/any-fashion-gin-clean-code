package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"golang.org/x/net/context"
)

type CouponRepository interface {
	AddCoupon(ctx context.Context, coupon req.Coupon) error
	DeleteCoupon(ctx context.Context, Couponid int) error
}
