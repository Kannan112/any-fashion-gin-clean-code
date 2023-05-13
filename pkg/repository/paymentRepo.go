package repository

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type PaymentDataBase struct {
	DB gorm.DB
}

func NewPaymentRepo(DB gorm.DB) interfaces.PaymentRepo {
	return PaymentDataBase{DB}
}

func (c PaymentDataBase) SavePaymentMethod(payment req.PaymentReq) error {
	query := `INSERT INTO payment_methods(payment_type,block_status,maximum_amount,created_at,updated_at) VALUES($1,$2,$3,NOW())`
	err := c.DB.Exec(query, payment.PaymentType, payment.BlockStatus, payment.MaximumAmount).Error
	return err
}

//update the pament mathod

func (c PaymentDataBase) UpdatePaymentMethod(pamyment req.PaymentReq) error {
	query := `UPDATE payment_method SET(payment_type=$1,block_status=$2,maximum_amount=$3,update_at=NOW())`
	err := c.DB.Exec(query, pamyment.PaymentType, pamyment.BlockStatus, pamyment.MaximumAmount).Error
	return err
}
