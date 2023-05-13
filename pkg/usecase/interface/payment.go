package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
)

type PaymentUsecases interface {
	SavePaymentMethod(payment req.PaymentReq) error
	UpdatePaymentMethod(Paymen req.PaymentReq)error
	
}
