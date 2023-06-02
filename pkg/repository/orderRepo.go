package repository

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB}
}

func (c *OrderDatabase) OrderAll(userId int) (domain.Order, error) {
	//get the cartid and userid and total of the cart
	var dom domain.Order
	tx := c.DB.Begin()
	var cart domain.Carts
	query := `SELECT * FROM carts WHERE users_id=$1`
	err := tx.Raw(query, userId).Scan(&cart).Error
	if err != nil {
		tx.Rollback()
		return dom, err
	}
	fmt.Println("cartid", cart.Id)
	if cart.Total == 0 {
		setTotal := `UPDATE carts SET total=carts.sub_total`
		err := tx.Exec(setTotal).Error
		if err != nil {
			tx.Rollback()
			return dom, err
		}
		cart.Total = cart.Sub_total
	}
	if cart.Sub_total == 0 {
		tx.Rollback()
		return dom, fmt.Errorf("NO ITEM IN CART")
	}
	//FIND THE DEFAULT ADDRESS OF THE USER
	var addressId int
	address := `SELECT id FROM addresses WHERE users_id=$1 AND is_default=true`
	err = tx.Raw(address, userId).Scan(&addressId).Error
	if err != nil {
		tx.Rollback()
		return dom, err
	}
	if addressId == 0 {
		tx.Rollback()
		return dom, fmt.Errorf("add address")
	}
	var order domain.Orders
	insetOrder := `INSERT INTO orders (users_id,order_time,address_id,order_total,order_status)
		VALUES($1,NOW(),$2,$3,'order placed') RETURNING *`
	err = tx.Raw(insetOrder, userId, addressId, cart.Total).Scan(&order).Error
	if err != nil {
		tx.Rollback()
		return dom, err
	}
	//GET CART ITEMS details of the user
	var cartItems []req.CartItems
	cartDetails := `select ci.product_item_id, ci.quantity ,pi.price,pi.qnty_in_stock FROM cart_items ci JOIN product_items pi on ci.product_item_id=pi.id where ci.carts_id=$1`
	err = tx.Raw(cartDetails, cart.Id).Scan(&cartItems).Error
	if err != nil {
		tx.Rollback()
		return dom, err
	}
	//Add the items in the cart into the orderitems one by one
	for _, items := range cartItems {
		fmt.Println("quantity", items)
		//check whether the item is available
		if items.Quantity > items.QntyInStock {
			return dom, fmt.Errorf("out of stock")
		}
		insetOrderItems := `INSERT INTO order_items (orders_id,product_item_id,quantity,price) VALUES($1,$2,$3,$4)`
		err = tx.Exec(insetOrderItems, order.Id, items.ProductItemId, items.Quantity, items.Price).Error

		if err != nil {
			tx.Rollback()
			return dom, err
		}
	}
	//Update the cart total
	updateCart := `UPDATE carts SET total=0,sub_total=0 WHERE users_id=?`
	err = tx.Exec(updateCart, userId).Error
	if err != nil {
		tx.Rollback()
		return dom, err
	}
	//Remove the items from the cart_items
	for _, items := range cartItems {
		removeCartItems := `DELETE FROM cart_items WHERE carts_id =$1 AND product_item_id=$2`
		err = tx.Exec(removeCartItems, cart.Id, items.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return dom, err
		}
	}
	//Reduce the product qty in stock details
	for _, items := range cartItems {
		updateQty := `UPDATE product_items SET qnty_in_stock=product_items.qnty_in_stock-$1 WHERE id=$2`
		err = tx.Exec(updateQty, items.Quantity, items.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return dom, err
		}
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return dom, err
	}
	return dom, nil

}
func (c *OrderDatabase) UserCancelOrder(orderId, userId int) (float32, error) {
	tx := c.DB.Begin()
	//collect order details
	var OrderDetails []req.CartItems
	CollectDetails := `SELECT product_item_id,quantity,price from order_items where orders_id=$1`
	err := tx.Raw(CollectDetails, orderId).Scan(&OrderDetails).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if len(OrderDetails) == 0 {
		return 0, fmt.Errorf("no order found with this id")
	}
	for _, items := range OrderDetails {
		updateProductItems := `UPDATE product_items SET qnty_in_stock=qnty_in_stock+$1 WHERE id=$2`
		err := tx.Exec(updateProductItems, items.Quantity, items.ProductItemId).Error
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		//remove order from order_items
		Remove := `DELETE from order_items WHERE orders_id=$1`
		err = tx.Exec(Remove, orderId).Error
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	updateOrder := `UPDATE orders SET order_status='order cancelled' WHERE id=$1 AND users_id=$2`
	err = tx.Exec(updateOrder, orderId, userId).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var price float32
	RefundAmount := `SELECT order_total FROM orders WHERE id=$1`
	err = tx.Raw(RefundAmount, orderId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return price, err
}
func (c *OrderDatabase) ListAllOrders(userId int) ([]domain.Order, error) {
	var order []domain.Order
	query := `SELECT *
	FROM orders	WHERE users_id = $1;
	`
	err := c.DB.Raw(query, userId).Scan(&order).Error
	return order, err

}
func (c *OrderDatabase) ListAllOrdersByStatus(userId, status int) ([]domain.Order, error) {
	var order []domain.Order
	query := `SELECT * FROM orders WHERE users_id=$1`
	err := c.DB.Raw(query, userId).Scan(&order).Error
	return order, err
}

func (c *OrderDatabase) OrderDetails(ctx context.Context, orderId uint, userId uint) ([]res.UserOrder, error) {
	var UserOrderDetails []res.UserOrder
	query := `  select * from orders JOIN users on orders.users_id=users.Id JOIN addresses on addresses.users_id=users.Id WHERE orders.id=$1 AND users.id=$2`
	err := c.DB.Raw(query, orderId, userId).Scan(&UserOrderDetails).Error
	return UserOrderDetails, err
}
