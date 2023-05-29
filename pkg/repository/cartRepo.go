package repository

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type CartDataBase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDataBase{DB}
}

func (c *CartDataBase) FindCart(ctx context.Context, userId int) (domain.Carts, error) {

	var cart domain.Carts
	query := `SELECT * FROM carts WHERE users_id = $1`
	err := c.DB.Raw(query, userId).Scan(&cart).Error
	fmt.Println("cart", cart)
	return cart, err
}

// Create cart
func (c *CartDataBase) CreateCart(id int) error {
	query := `INSERT INTO carts(users_id,sub_total,total)VALUES($1,0,0)`
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
	//check the product exist in the cart_item
	var CartitemID int
	cartItemCheck := `SELECT id FROM cart_items WHERE carts_id=$1 AND product_item_id=$2 LIMIT 1`
	err = c.DB.Raw(cartItemCheck, cartId, productId).Scan(&CartitemID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if CartitemID == 0 {
		addtoCart := `INSERT INTO cart_items(carts_id, product_item_id, quantity) VALUES($1, $2, 1)`
		err = tx.Exec(addtoCart, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		updateCart := `UPDATE cart_items SET quantity=cart_items.quantity+1 WHERE id=$1`
		err = tx.Exec(updateCart, CartitemID).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	var Price int
	findPrice := `SELECT price FROM product_items WHERE id=$1`
	err = tx.Raw(findPrice, productId).Scan(&Price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//update subtotal in cart table
	var subtotal int
	updateTotal := `UPDATE carts SET sub_total=sub_total+$1 where users_id=$2 RETURNING sub_total`
	err = tx.Raw(updateTotal, Price, userId).Scan(&subtotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//check any coupon is present in the cart
	var couponId int
	CheckCouponId := `SELECT coupon_id FROM carts WHERE users_id=$1`
	err = tx.Raw(CheckCouponId, userId).Scan(&couponId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if couponId != 0 {
		//find coupon details
		var couponDetails domain.Coupon
		CouponTable := `SELECT * FROM coupons WHERE id=$1`
		err := tx.Raw(CouponTable, couponId).Scan(&couponDetails).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		//apply the coupon to the total
		discountAmount := (subtotal / 100) * int(couponDetails.DiscountPercent)
		if discountAmount > int(couponDetails.DiscountMaximumAmount) {
			discountAmount = int(couponDetails.DiscountMaximumAmount)
		}
		updateCart := `UPDATE carts SET total=$1 WHERE id =$2`
		err = tx.Exec(updateCart, subtotal-discountAmount, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		updateTotal := `UPDATE carts SET total=$1 where id=$2`
		err := tx.Exec(updateTotal, subtotal, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

// Remove items from cart
func (c *CartDataBase) RemoveFromCart(userId int, ProductItemId int) error {
	tx := c.DB.Begin()
	//find the cart id
	var cartID int
	query2 := `SELECT id FROM carts WHERE users_id=$1`
	err := tx.Raw(query2, userId).Scan(&cartID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//find the qty of product in cart items
	var qty int
	findQTY := `SELECT quantity FROM cart_items WHERE carts_id=$1 AND product_item_id=$2`
	err = tx.Raw(findQTY, cartID, ProductItemId).Scan(&qty).Error
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
		err = tx.Exec(delete, cartID, ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		//If Qty is more the one QTY reduce the QTY
	} else {
		update := `UPDATE cart_items SET quantity=cart_items.quantity-1 WHERE carts_id=$1 AND product_item_id=$2`
		err = tx.Exec(update, cartID, ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//find the price of product item
	var price int
	productPrice := `SELECT price FROM product_items WHERE id =$1`
	err = tx.Raw(productPrice, ProductItemId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//Update the subtotal in cart table
	var subtotal int
	UpdateSubTotal := `UPDATE carts SET sub_total=sub_total-$1 WHERE users_id=$2 RETURNING sub_total`
	err = tx.Raw(UpdateSubTotal, price, userId).Scan(&subtotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

// Lish Cart
func (c *CartDataBase) ListCart(userId int) ([]domain.Cart, error) {
	var list []domain.Cart

	query := `SELECT * FROM carts WHERE users_id=$1`
	err := c.DB.Raw(query, userId).Scan(&list).Error
	return list, err
}
