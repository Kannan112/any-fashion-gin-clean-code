package repository

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type PaymentDataBase struct {
	DB *gorm.DB
}

func NewPaymentRepo(DB *gorm.DB) interfaces.PaymentRepo {
	return &PaymentDataBase{DB}
}

func (c PaymentDataBase) SavePaymentMethod(payment req.PaymentReq) error {
	query := `INSERT INTO payment_methods(payment_type,block_status,maximum_amount,created_at,updated_at) VALUES($1,$2,$3,NOW(),0)`
	err := c.DB.Exec(query, payment.PaymentType, payment.BlockStatus, payment.MaximumAmount).Error
	return err
}

//update the pament mathod

func (c PaymentDataBase) UpdatePaymentMethod(id int, pamyment req.PaymentReq) error {
	//check the id

	chech := `SELECT id FROM payment_methods WHERE id=$1`
	err := c.DB.Exec(chech, id).Error
	if err != nil {
		return err
	}

	query := `UPDATE payment_methods SET payment_type=$1, block_status=$2, maximum_amount=$3, updated_at=NOW() WHERE id=$4`
	err = c.DB.Exec(query, pamyment.PaymentType, pamyment.BlockStatus, pamyment.MaximumAmount, id).Error
	return err
}
