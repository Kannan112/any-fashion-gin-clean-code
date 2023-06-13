package repository

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type WishListDataBase struct {
	DB *gorm.DB
}

func NewWishlistRepository(DB *gorm.DB) interfaces.WishListRepo {
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
		return fmt.Errorf("item not found")
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
	insert := `INSERT INTO wish_lists(users_id,item_id)VALUES($1,$2)`
	err = tx.Exec(insert, id, itemId).Error
	if err != nil {
		fmt.Println("Test Me")
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

func (c *WishListDataBase) ListAllWishlist(ctx context.Context, userId int, pagenation req.Pagenation) ([]res.ProductItem, error) {
	var wishlists []res.ProductItem
	var err error

	if pagenation.Count == 0 || pagenation.Page == 0 {
		err = c.DB.Table("wish_lists").
			Select("pi.id, wish_lists.item_id, pi.sku, pi.qnty_in_stock, pi.color, pi.gender, pi.material, pi.size, pi.model, pi.price, p.product_name, p.description, p.brand, c.name").
			Joins("JOIN product_items pi ON wish_lists.item_id = pi.id").
			Joins("JOIN products p ON pi.product_id = p.id").
			Joins("JOIN categories c ON p.category_id = c.id").
			Where("wish_lists.users_id = ?", userId).
			Scan(&wishlists).Error
	} else {
		limit := pagenation.Count
		offset := (pagenation.Page - 1) * limit
		var check bool
		query := `SELECT EXISTS (SELECT 1 FROM wish_lists)`
		err = c.DB.Raw(query).Scan(&check).Error
		if err != nil {
			return nil, err
		}
		if !check {
			return nil, fmt.Errorf("wish list is empty")
		}
		query2 := `SELECT pi.id, w.item_id, pi.sku, pi.qnty_in_stock, pi.color, pi.gender, pi.material, pi.size, pi.model, pi.price, p.product_name, p.description, p.brand, c.name 
			FROM wish_lists w
			JOIN product_items pi ON w.item_id = pi.id
			JOIN products p ON pi.product_id = p.id
			JOIN categories c ON p.category_id = c.id
			WHERE w.users_id = $1
			OFFSET $2
			LIMIT $3;`
		err = c.DB.Raw(query2, userId, offset, limit).Scan(&wishlists).Error
	}

	return wishlists, err
}

// func (c *WishListDataBase) OrderWishlist(ctx context.Context, userId int) error {
// 	var check bool
// 	tx := c.DB.Begin()
// 	checkingWishlist := `SELECT EXISTS(SELECT * FROM wishlist where users_id=$1)`
// 	err := tx.Raw(checkingWishlist, userId).Scan(&check).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	if !check{
// 		tx.Rollback()
// 		return fmt.Errorf("Your wishlist is empty")
// 	}
// }
