package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
)

type PaymentRepo interface {
	SavePaymentMethod(payment req.PaymentReq) error
	UpdatePaymentMethod(id int, payment req.PaymentReq) error
}