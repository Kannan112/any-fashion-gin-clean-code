package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
)

type PaymentUsecases interface {
	SavePaymentMethod(payment req.PaymentReq) error
	UpdatePaymentMethod(id int, Paymen req.PaymentReq) error
}
