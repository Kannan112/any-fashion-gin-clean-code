package db

import (
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	"gorm.io/gorm"
)

func SavePaymentMethods(db *gorm.DB) error {
	Payment := []domain.PaymentMethod{
		{
			Name:          domain.RazopayPayment,
			MaximumAmount: domain.RazorayMaxiumAmount,
		},
		{
			Name:          domain.CodePayment,
			MaximumAmount: domain.CodMaxiumAmount,
		},
		{
			Name:          domain.StripePayment,
			MaximumAmount: domain.StripeMaxiumAmount,
		},
	}
	var (
		searchQuery = "select case when id != 0 THEN 'T' ELSE 'F' END as exist FROM payment_methods where name=$1 "
		insertQuery = `INSERT INTO payment_methods (name, maximum_amount) VALUES ($1, $2)`
		exist       bool
		err         error
	)

	for _, paymentMethod := range Payment {
		err = db.Raw(searchQuery, paymentMethod.Name).Scan(&exist).Error
		if err != nil {
			return fmt.Errorf("failed to check payment methods already exist %w", err)
		}
		if !exist {
			err = db.Exec(insertQuery, paymentMethod.Name, paymentMethod.MaximumAmount).Error
			if err != nil {
				return fmt.Errorf("failed to save payment method %w", err)
			}
		}
		exist = false
	}
	return nil
}
