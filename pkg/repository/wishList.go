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

func (c *WishListDataBase) AddToWishlist(id, productId int) error {
	tx := c.DB.Begin()
	var checkPresence bool
	query := `SELECT EXISTS (SELECT 1 FROM wish_lists WHERE users_id = $1 AND product_id = $2);`
	err := c.DB.Raw(query, id, productId).Scan(&checkPresence).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if checkPresence {
		tx.Rollback()
		return fmt.Errorf("the same product is already added to wishlist")
	}
	insert := `INSERT INTO wish_lists(product_id,users_id)VALUES($1,$2)`
	err = tx.Exec(insert, productId, id).Error
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

func (c *WishListDataBase) RemoveFromWishlist(ctx context.Context, userid, productid int) error {
	tx := c.DB.Begin()
	var check bool
	query := `SELECT EXISTS (SELECT 1 FROM wish_lists WHERE users_id = $1 AND product_id = $2)`
	err := tx.Raw(query, userid, productid).Scan(&check).Error
	if err != nil {
		fmt.Println("errfor1", err)
		tx.Rollback()
		return err
	}
	if !check {
		tx.Rollback()
		return fmt.Errorf("the item is not present in the wishlist")
	}

	query2 := `DELETE FROM wish_lists WHERE users_id=$1 AND product_id = $2`
	err = tx.Exec(query2, userid, productid).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = c.DB.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
