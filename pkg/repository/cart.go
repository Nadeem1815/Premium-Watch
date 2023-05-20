package repository

import (
	"context"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	"gorm.io/gorm"
)

type CartDataBase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDataBase{DB}
}

func (cr *CartDataBase) AddToCart(ctx context.Context, userID string, productID int) (domain.CartItems, error) {
	// Begi transaction
	tx := cr.DB.Begin()
	// finding cart id corresponding the user
	var cartID int
	fmt.Println(userID)
	findCartId := `SELECT id FROM carts WHERE user_id=? LIMIT 1 `
	err := tx.Raw(findCartId, userID).Scan(&cartID).Error

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err

	}
	if cartID == 0 {

		// if user has no cart,creating one
		err := tx.Raw("INSERT INTO carts(user_id,sub_total,total)VALUES($1,0,0)RETURNING id", userID).Scan(&cartID).Error

		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}
	}

	// checking if product is already present in the cart
	var cartItem domain.CartItems
	err = tx.Raw("SELECT id,quantity FROM cart_items WHERE cart_id=$1 AND product_id=$2 LIMIT 1", cartID, productID).Scan(&cartItem).Error

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err

	}
	// if item is not present in the cart
	if cartItem.ID == 0 {
		err = tx.Raw("INSERT INTO cart_items(cart_id,product_id,quantity)VALUES($1,$2,1)RETURNING*;", cartID, productID).Scan(&cartItem).Error
		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}

	} else {
		// if item is already present in the cart
		err = tx.Raw("UPDATE cart_items SET quantity=$1 WHERE id=$2 RETURNING*;", cartItem.Quantity+1, cartItem.ID).Scan(&cartItem).Error
		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err
		}
	}

	// update subtotal in cart table
	// product_id is know ,cart_id is known,quantity is known
	// fetch current subtotal from cart table
	var itemPrice float64
	err = tx.Raw("SELECT price FROM products WHERE id=$1", productID).Scan(&itemPrice).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	type totalsprice struct {
		currentSubTotal float64
		total           float64
	}
	var totals totalsprice
	// fetch current subtotal from  cart table
	err = tx.Raw("SELECT sub_total, total FROM carts WHERE id=$1", productID).Scan(&totals).Error
	// err=tx.Raw("SELECT total FROM carts WHERE id=$1",productID).Scan(&total).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	// add price of new product item of the current subtotal and product it in the cart tableconst
	newSubTotal := totals.currentSubTotal + itemPrice
	newTotal := totals.total + itemPrice
	err = tx.Exec("UPDATE carts SET sub_total=$1,total=$2 WHERE user_id=$3", newSubTotal, newTotal, userID).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	return cartItem, nil

}
