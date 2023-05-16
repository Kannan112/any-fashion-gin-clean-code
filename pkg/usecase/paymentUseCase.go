package usecase

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type PaymentUsecase struct {
	paymentrepo interfaces.PaymentRepo
}

func NewPaymentUsecase(paymentrepo interfaces.PaymentRepo) services.PaymentUsecases {
	return &PaymentUsecase{
		paymentrepo: paymentrepo,
	}
}

func (c *PaymentUsecase) SavePaymentMethod(payment req.PaymentReq) error {
	err := c.paymentrepo.SavePaymentMethod(payment)
	return err
}

func (c *PaymentUsecase) UpdatePaymentMethod(id int, Payment req.PaymentReq) error {
	err := c.paymentrepo.UpdatePaymentMethod(id, Payment)
	return err
}
