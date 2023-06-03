package repository

import (
	"context"
	"errors"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
)

func (c *ProductDataBase) SaveOffer(ctx context.Context, offerdetails req.OfferTable) error {
	// Check if the product_id exists in the product table
	var count bool
	// err := c.DB.Model(&domain.Product{}).Where("id = ?", offerdetails.ProductId).Scan(&count).Error
	// if err != nil {
	// 	return err
	// }

	query := `SELECT EXISTS(SELECT * FROM products WHERE id=$1)`
	err := c.DB.Raw(query, offerdetails.ProductId).Scan(&count).Error
	if err != nil {
		return err
	}
	if !count {
		return errors.New("product not found")
	}

	query2 := `INSERT INTO offer_tables(product_id, discount, start_date, end_date,discription) VALUES ($1, $2, NOW(), $3, $4)`
	err = c.DB.Exec(query2, offerdetails.ProductId, offerdetails.Discount, offerdetails.EndDate, offerdetails.Discription).Error
	if err != nil {
		return err
	}
	return nil
}
