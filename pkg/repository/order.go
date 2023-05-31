package repository

import (
	"context"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepoitory {
	return &OrderDatabase{DB}
}

func (cr *OrderDatabase) BuyAll(ctx context.Context, body model.PlaceOrder, userID string) (domain.Order, error) {
	tx := cr.DB.Begin()
	var cartDetails struct {
		ID    int
		Total float64
	}
	findCartQuery := `SELECT id,total FROM carts WHERE user_id=$1`
	err := tx.Raw(findCartQuery, userID).Scan(&cartDetails).Error
	if cartDetails.ID == 0 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("no items in cart")

	}
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}
	var cartItems []domain.CartItems
	fetchCartItems := `SELECT *FROM cart_items WHERE cart_id=$1;`
	err = tx.Raw(fetchCartItems, cartDetails.ID).Scan(&cartItems).Error

	if len(cartItems) == 0 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("cart is empty")

	}

	var createdOrder domain.Order
	// order createing
	orderquery := `INSERT INTO orders(user_id,order_date,shipping_address_id,order_total,order_status_id)
				 VALUES($1,NOW(),$2,$3,1) RETURNING *;`
	err = tx.Raw(orderquery, userID, body.ShippingAddressID, cartDetails.Total).Scan(&createdOrder).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err

	}

	// update cartItems table
	deletecartItemQuery := `DELETE FROM cart_items WHERE cart_id=$1`

	err = tx.Exec(deletecartItemQuery, cartDetails.ID).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err

	}
	createOrderItemQuery := `INSERT INTO order_items(product_id,order_id,quantity,price)VALUES($1,$2,$3,$4)`

	for i := range cartItems {
		// fetch the stock
		var productDetails struct {
			Stock int
			Price float64
		}
		fetchProducts := `SELECT stock,price FROM products WHERE id=$1;`
		err = tx.Raw(fetchProducts, cartItems[i].ProductID).Scan(&productDetails).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}
		// if products is outofstock
		if productDetails.Stock < int(cartItems[i].Quantity) {
			tx.Rollback()
			return domain.Order{}, err
		}
		// creating order items
		productTotal := productDetails.Price * float64(cartItems[i].Quantity)
		err = tx.Exec(createOrderItemQuery, cartItems[i].ProductID, createdOrder.ID, cartItems[i].Quantity, productTotal).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err

		}

		// decrease  the stock of product
		decreaseQuery := `UPDATE products SET stock=stock -$1 WHERE id=$2`
		err = tx.Exec(decreaseQuery, cartItems[i].Quantity, cartItems[i].ProductID).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}
	}
	tx.Commit()
	return createdOrder, nil
}

func (cr *OrderDatabase) CancelOrder(ctx context.Context, orderID int, UserID string) (domain.Order, error) {
	tx := cr.DB.Begin()
	// find order details. if order is pending user can cancel if order is not pending user cant cancel

	var orderStatusId int
	viewStatusQuery := `SELECT order_status_id FROM orders WHERE user_id=$1 AND id=$2`
	err := tx.Raw(viewStatusQuery, UserID, orderID).Scan(&orderStatusId).Error
	if err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}
	if orderStatusId == 0 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("no such order found")
	}

	// if order is pending
	if orderStatusId == 1 {
		var cancelledOrder domain.Order
		cancelQuery := `UPDATE orders SET order_status_id=2 WHERE user_id=$1 AND id=$2 RETURNING *;`
		err := tx.Raw(cancelQuery, UserID, orderID).Scan(&cancelledOrder).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}

		// increase the product Item table
		var orderItem []domain.OrderItem
		findOrderItemsQuery := `SELECT *FROM order_items WHERE order_id=$1`
		err = tx.Raw(findOrderItemsQuery, orderID).Scan(&orderItem).Error
		if err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}
		qntyUpdateQuery := `UPDATE products SET stock=stock+$1 WHERE id=$2`
		for i := range orderItem {
			err := tx.Exec(qntyUpdateQuery, orderItem[i].Quantity, orderItem[i].ProductID).Error
			if err != nil {
				tx.Rollback()
				return domain.Order{}, err
			}
		}
		tx.Commit()
		return cancelledOrder, nil
	}

	// if order already cancelled
	if orderStatusId == 2 {
		tx.Rollback()
		return domain.Order{}, fmt.Errorf("order already cancelled")

	}
	tx.Rollback()
	return domain.Order{}, fmt.Errorf("order processed ,cannot cancelled")

}

func (cr *OrderDatabase) UpdateOrder(ctx context.Context, orderInfo model.UpdateOrder) (domain.Order, error) {
	var updatedOrder domain.Order

	updateQuery := `UPDATE orders SET order_status_id=$1,delivery_status_id=$2,delivery_updated_at=NOW() WHERE id=$3 RETURNING*;`
	err := cr.DB.Raw(updateQuery, orderInfo.OrderStatusID, orderInfo.DeliveryStatusId, orderInfo.OrderID).Scan(&updatedOrder).Error
	if err != nil {
		return domain.Order{}, err

	}
	if updatedOrder.ID == 0 {
		return domain.Order{}, fmt.Errorf("no order")

	}
	return updatedOrder, nil
}
