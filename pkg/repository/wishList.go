package repository

import (
	"context"
	"fmt"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type WishListDataBase struct {
	DB *gorm.DB
}

func NewWhishlistRepository(DB *gorm.DB) interfaces.WishListRepo {
	return &WishListDataBase{
		DB: DB,
	}
}

func (c *WishListDataBase) AddToWishlist(ctx context.Context, id, productId int) error {
	tx := c.DB.Begin()
	var checkPresence bool
	query := `SELECT EXIST(SELECT ID FROM wish_lists WHERE users_id=$1 AND product_id=$2)`
	err := c.DB.Raw(query, id, productId).Scan(&checkPresence).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if checkPresence {
		tx.Rollback()
		return fmt.Errorf("the same product is already added to wishlist")
	}
	insert := `INSERT INTO wish_list(users_id,product_id)VALUES($1,$2)`
	err = tx.Exec(insert, id, productId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// func RemoveFromWishlist(ctx context.Context,id int )error[

// ]
