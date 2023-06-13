package repository

import (
	"context"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"gorm.io/gorm"
)

type CartDataBase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDataBase{DB}
}

func (cr *CartDataBase) AddToCart(ctx context.Context, userID string, productID int) (domain.CartItems, error) {
	// Begin transaction
	tx := cr.DB.Begin()
	// finding cart id corresponding the user
	var cartID int

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
	fmt.Println(cartID)

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

	var total, sub_total float64
	// fetch current subtotal from  cart table

	err = tx.Raw("SELECT sub_total  FROM carts WHERE id=$1", cartID).Scan(&sub_total).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err

	}

	err = tx.Raw("SELECT total FROM carts WHERE id=$1", cartID).Scan(&total).Error

	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}

	// add price of new product item of the current subtotal and product it in the cart tableconst
	newSubTotal := sub_total + itemPrice
	newTotal := total + itemPrice
	err = tx.Exec("UPDATE carts SET sub_total=$1,total=$2 WHERE user_id=$3", newSubTotal, newTotal, userID).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}

	// check if the cart has a coupon
	var coupontID int
	err = tx.Raw("SELECT COALESCE(coupon_id,0)FROM carts WHERE user_id=$1", userID).Scan(&coupontID).Error
	if err != nil {
		tx.Rollback()
		return domain.CartItems{}, err

	}
	// if cart has coupon
	if coupontID != 0 {
		// fetch coupon details
		var couponInfo domain.Coupon
		err = tx.Raw("SELECT *FROM coupons WHERE id=$1", coupontID).Scan(&couponInfo).Error
		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err

		}
		discount := newSubTotal * (couponInfo.DiscountMaxAmount / 100)
		if discount > couponInfo.DiscountMaxAmount {
			discount = couponInfo.DiscountMaxAmount

		}
		updatedTotal := newTotal - discount
		// update cart table
		err = tx.Exec("UPDATED carts SET discount=&1,total=$2 WHERE user_id=$3", discount, updatedTotal, userID).Error
		if err != nil {
			tx.Rollback()
			return domain.CartItems{}, err

		}
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.CartItems{}, err
	}
	return cartItem, nil

}

func (cr *CartDataBase) RemoveTOCart(ctx context.Context, userID string, productId int) error {
	tx := cr.DB.Begin()
	// find cart_id from cart table
	var cartID int
	err := tx.Raw("SELECT id FROM carts WHERE user_id=$1", userID).Scan(&cartID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// find quantity
	var quantity int
	err = tx.Raw("SELECT quantity FROM cart_items WHERE cart_id=$1 AND product_id=$2", cartID, productId).Scan(&quantity).Error
	if err != nil {
		tx.Rollback()
		return err

	}

	// if quantity is 1 delete the row
	if quantity == 0 {
		tx.Rollback()
		return fmt.Errorf("nothing to remove")
	} else if quantity == 1 {
		err := tx.Exec("DELETE  FORM cart_items WHERE cart_id=&1 AND product_id=&2", cartID, productId).Error
		if err != nil {
			tx.Rollback()
			return err

		}
	} else {
		err := tx.Exec("UPDATE cart_items SET quantity=cart_items.quantity-$1 WHERE cart_id=$2 AND product_id=$3", 1, cartID, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	// fetch price from product table
	var itemPrice float64

	err = tx.Raw("SELECT price FROM products WHERE id=$1", productId).Scan(&itemPrice).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	var newSubTotal float64
	err = tx.Raw("UPDATE carts SET sub_total=sub_total -$1,total =total-$2 WHERE id=$3 RETURNING sub_total;", itemPrice, itemPrice, cartID).Scan(&newSubTotal).Error
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

func (cr *CartDataBase) ViewCart(ctx context.Context, userID string) (model.ViewCart, error) {
	tx := cr.DB.Begin()

	//  find cart_id from cart tables
	var cartDetails struct {
		ID       int
		CouponID int
		SubTotal float64
		Discount float64
		Total    float64
	}
	err := tx.Raw("SELECT id,coupon_id,sub_total,total FROM carts WHERE user_id=$1", userID).Scan(&cartDetails).Error
	if err != nil {
		tx.Rollback()
		return model.ViewCart{}, err
	}
	// fmt.Printf("%+v", cartDetails)
	var items []model.DisplayCart

	selectItems := ` select p.id,p.name, p.price, p.brand, p.colour,p.product_image,p.sku,c.quantity,c.item_total as total from products p JOIN cart_items c on c.product_id=p.id where c.cart_id=$1`
	err = tx.Raw(selectItems, cartDetails.ID).Scan(&items).Error
	if err != nil {
		tx.Rollback()
		return model.ViewCart{}, err
	}
	fmt.Println(items)
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return model.ViewCart{}, err
	}

	var finalCart model.ViewCart
	finalCart.CouponID = cartDetails.CouponID
	finalCart.SubTotal = cartDetails.SubTotal
	finalCart.Discount = cartDetails.Discount
	finalCart.Total = cartDetails.Total

	finalCart.CartItmes = items

	return finalCart, nil
}

func (cr *CartDataBase) AddCouponToCart(ctx context.Context, userID string, couponID int) (model.ViewCart, error) {
	// fetch couponDetail
	var couponInfo domain.Coupon
	fetchCoupon := `SELECT *FROM coupons WHERE id=$1`
	err := cr.DB.Raw(fetchCoupon, couponID).Scan(&couponInfo).Error
	if err != nil {
		return model.ViewCart{}, err

	}
	if couponInfo.ID == 0 {
		return model.ViewCart{}, fmt.Errorf("no coupon found")

	}
	// fetch cart details
	var cartInfo domain.Cart
	fetchCart := `SELECT *FROM carts WHERE user_id=$1`
	err = cr.DB.Raw(fetchCart, userID).Scan(&cartInfo).Error
	if err != nil {
		return model.ViewCart{}, err
	}
	if cartInfo.ID == 0 {
		return model.ViewCart{}, fmt.Errorf("your cart is empty can't add coupon")

	}
	// check this coupon for can apply for this cart total a mount
	if cartInfo.SubTotal < couponInfo.MinOrderValue {
		return model.ViewCart{}, fmt.Errorf("this coupon cant apply for this cart amount")

	}
	//  calculate discount amount
	discount := cartInfo.SubTotal * (couponInfo.DiscountPercent / 100)
	if discount > couponInfo.DiscountMaxAmount {
		discount = couponInfo.DiscountMaxAmount

	}
	total := cartInfo.SubTotal - discount

	// update cart
	updateCartQuery := `UPDATE carts SET coupon_id=$1, discount=$2,total=$3 WHERE user_id = $4`
	err = cr.DB.Exec(updateCartQuery, couponID, discount, total, userID).Error
	if err != nil {
		return model.ViewCart{}, err

	}
	cart, err := cr.ViewCart(ctx, userID)
	if err != nil {
		return model.ViewCart{}, err

	}
	return cart, nil
}

