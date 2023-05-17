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

func (c *WishListDataBase) AddToWishlist(id, itemId int) error {
	tx := c.DB.Begin()
	var checkPresence bool
	var Exists bool
	find := `SELECT EXISTS(SELECT id FROM product_items WHERE id=$1)`
	err := tx.Raw(find, itemId).Scan(&Exists).Error
	if err != nil {
		return err
	}
	if !Exists {
		return fmt.Errorf("Item not found")
	}

	query := `SELECT EXISTS (SELECT 1 FROM wish_lists WHERE users_id = $1 AND item_id = $2);`
	err = c.DB.Raw(query, id, itemId).Scan(&checkPresence).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if checkPresence {
		tx.Rollback()
		return fmt.Errorf("the same product is already added to wishlist")
	}
	insert := `INSERT INTO wish_lists(item_id,users_id)VALUES($1,$2)`
	err = tx.Exec(insert, itemId, id).Error
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

func (c *WishListDataBase) RemoveFromWishlist(ctx context.Context, userid, itemId int) error {
	tx := c.DB.Begin()
	var check bool

	query := `SELECT EXISTS (SELECT 1 FROM wish_lists WHERE users_id = $1 AND item_id = $2)`
	err := tx.Raw(query, userid, itemId).Scan(&check).Error
	if err != nil {

		tx.Rollback()
		return err
	}
	if !check {
		tx.Rollback()
		return fmt.Errorf("the item is not present in the wishlist")
	}

	query2 := `DELETE FROM wish_lists WHERE users_id=$1 AND item_id = $2`
	err = tx.Exec(query2, userid, itemId).Error
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

// func (c *WishListDataBase) ListAllWishlist(ctx context.Context, userId int) (res.Wishlist, error) {
// 	var check bool
// var wishlist res.Wishlist
// 	query := `SELECT EXISTS(SELECT 1 FROM wish_lists WHERE users_id=$1)`
// 	err := c.DB.Exec(query, userId).Scan(&check).Error
// 	if err != nil {
// 		return wishlist, err
// 	}
// 	if !check {
// 		return wishlist, fmt.Errorf("the item is not present in the wishlist")
// 	}
// 	query1 := `SELECT product_id FROM `

// }
