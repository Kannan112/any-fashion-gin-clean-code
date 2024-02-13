package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type PaymentUsecases interface {
	ListPaymentMethod(ctx context.Context) ([]domain.PaymentMethod, error)
	UpdatePaymentMethod(id int, Paymen req.PaymentReq) error

	// razorpay
	RazorPayCheckout(ctx context.Context, userId int) (res.RazorpayOrder, error)
	VerifyRazorPay(ctx context.Context, body req.RazorPayRequest) error

	// stripe
	MakeStripeOrder(ctx context.Context, userID uint) (res.StripeOrder, error)
	VerifyStripeOrder(ctx context.Context, stripePaymentID string) error
}
