package repository

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type PaymentDataBase struct {
	DB *gorm.DB
}

func NewPaymentRepo(DB *gorm.DB) interfaces.PaymentRepo {
	return &PaymentDataBase{DB}
}

func (c PaymentDataBase) ListPaymentMethod(ctx context.Context) ([]domain.PaymentMethod, error) {
	var data []domain.PaymentMethod
	ListPayment := `select * from payment_methods`
	if err := c.DB.Raw(ListPayment).Scan(&data).Error; err != nil {
		return data, err
	}
	return data, nil

}

//update the pament mathod

func (c PaymentDataBase) UpdatePaymentMethod(id int, pamyment req.PaymentReq) error {
	//check the id

	chech := `SELECT id FROM payment_methods WHERE id=$1`
	err := c.DB.Exec(chech, id).Error
	if err != nil {
		return err
	}

	query := `UPDATE payment_methods SET block_status=$1, maximum_amount=$2 WHERE id=$3`
	err = c.DB.Exec(query, pamyment.BlockStatus, pamyment.MaximumAmount, id).Error
	return err
}
