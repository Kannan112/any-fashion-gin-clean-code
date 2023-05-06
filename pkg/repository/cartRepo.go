package repository

import (
	"fmt"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type CartDataBase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDataBase{DB}
}

func (c *CartDataBase) CreateCart(id int) error {
	query := `INSERT INTO carts(user_id,total)VALUES($1,0)`
	err := c.DB.Exec(query, id).Error
	return err
}
func (c *CartDataBase) AddToCart(productId, userId int) error {
	tx := c.DB.Begin()

	var cartId int
	findcart := `SELECT id FROM carts WHERE users_id=$1`
	err := c.DB.Raw(findcart, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if cartId == 0 {
		createCart := `INSERT INTO carts (users_id,total) VALUES ($1,$2) RETURNING id`
		err = c.DB.Raw(createCart, userId, 0).Scan(&cartId).Error
		if err != nil {
			return err
		}
	}
	//check the product exist in the cart_item
	var CartitemID int
	cartItemCheck := `SELECT id FROM cart_items WHERE carts_id=$1 AND product_id=$2 LIMIT 1`
	err = c.DB.Raw(cartItemCheck, cartId, productId).Scan(&CartitemID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if CartitemID == 0 {
		addtoCart := `INSERT INTO cart_items(carts_id, product_id, quantity) VALUES($1, $2, 0) RETURNING id`
		err = c.DB.Raw(addtoCart, cartId, productId).Scan(&CartitemID).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		update := `UPDATE cart_items SET quantity=cart_items.quantity+1 WHERE id=$1`
		err = tx.Exec(update, CartitemID).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	var Price int
	findPrice := `SELECT price FROM products WHERE id=$1`
	err = tx.Raw(findPrice, productId).Scan(&Price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//update total in cart table
	var total int
	updateTotal := `UPDATE carts SET total=total+$1 WHERE users_id=$2 RETURNING total`
	err = tx.Raw(updateTotal, Price, userId).Scan(&total).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// func (c *CartDataBase) AddQuantity(ctx context.Context, cartItemID uint, qty uint) error {
// 	query := `UPDATE cart_items SET qty=$1 WHERE id=$2`
// 	if c.DB.Exec(query, qty, cartItemID).Error != nil {
// 		return fmt.Errorf("failed to add qty")
// 	}
// 	return nil
// }

func (c *CartDataBase) RemoveFromCart(userId int, ProductId int) error {
	tx := c.DB.Begin()
	//find the cart id
	var cartID int
	cartId := `SELECT id FROM carts WHERE user_id=$1`
	err := tx.Raw(cartId).Scan(&cartID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//find the qty of product in cart items
	var qty int
	findQTY := `SELECT quantity FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
	err = tx.Raw(findQTY, cartID, ProductId).Scan(&qty).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if qty == 0 {
		tx.Rollback()
		return fmt.Errorf("no items in the cart to remove")
	}
	//If qty is one DELETE item
	if qty == 1 {
		delete := `DELETE FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
		err = tx.Exec(delete, cartID, ProductId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		//If Qty is more the one QTY reduce the QTY
	} else {
		update := `UPDATE FROM cart_items SET quantity=cart_items.quantity-1 WHERE carts_id=$1 AND product_item_id=$2`
		err = tx.Exec(update, cartID, ProductId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//find the price of product item
	var price int
	productPrice := `SELECT price FROM product_items WHERE id =$1`
	err = tx.Exec(productPrice, ProductId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//Update the subtotal in cart table
	var subtotal int
	UpdateSubTotal := `UPDATE carts SET total=total-$1 WHERE user_id=$2 RETURNING sub_total`
	err = tx.Raw(UpdateSubTotal, price, userId).Scan(&subtotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}
