package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type PaymentRepo interface {
	ListPaymentMethod(ctx context.Context) ([]domain.PaymentMethod, error)
	UpdatePaymentMethod(id int, payment req.PaymentReq) error
}
