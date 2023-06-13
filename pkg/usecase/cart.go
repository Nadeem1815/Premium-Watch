package usecase

import (
	"context"
	"fmt"

	"time"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type CartUseCase struct {
	cartRepo    interfaces.CartRepository
	productRepo interfaces.ProductRepository
}

func NewCartUseCase(repo interfaces.CartRepository, product interfaces.ProductRepository) services.CartUseCase {
	return &CartUseCase{
		cartRepo:    repo,
		productRepo: product,
	}
}

func (cr *CartUseCase) AddToCart(ctx context.Context, userID string, productID int) (domain.CartItems, error) {
	cartItem, err := cr.cartRepo.AddToCart(ctx, userID, productID)
	return cartItem, err
}

func (cr *CartUseCase) RemoveTOCart(ctx context.Context, userID string, productId int) error {
	err := cr.cartRepo.RemoveTOCart(ctx, userID, productId)
	return err
}

func (cr *CartUseCase) ViewCart(ctx context.Context, userID string) (model.ViewCart, error) {
	viewCart, err := cr.cartRepo.ViewCart(ctx, userID)
	return viewCart, err
}

func (cr *CartUseCase) AddCouponToCart(ctx context.Context, userID string, couponID int) (model.ViewCart, error) {
	// check is coupon is already used
	isUsed, err := cr.productRepo.CouponUsed(ctx, userID, couponID)
	if err != nil {
		return model.ViewCart{}, err

	}
	if isUsed {
		return model.ViewCart{}, fmt.Errorf("coupon already is used")

	}
	// fetching coupon deatils
	couponDtls, err := cr.productRepo.ViewCouponById(ctx, couponID)
	if err != nil {
		return model.ViewCart{}, err

	}
	if couponDtls.ID == 0 {
		return model.ViewCart{}, fmt.Errorf("invalid coupon id")

	}
	// check coupon is valid
	currentTime := time.Now()

	if couponDtls.ValidTill.Before(currentTime) {

		return model.ViewCart{}, fmt.Errorf("coupon is expired")

	}
	// fetch cart totals
	cartInfo, err := cr.cartRepo.ViewCart(ctx, userID)
	if err != nil {
		return model.ViewCart{}, err

	}
	if cartInfo.SubTotal < couponDtls.MinOrderValue {
		return model.ViewCart{}, fmt.Errorf("this coupon cant apply for this cart amount")

	}
	// add coupon to cart
	couponAddcart, err := cr.cartRepo.AddCouponToCart(ctx, userID, couponID)
	// if err!=nil {
	// 	return model.ViewCart{},err

	// }
	return couponAddcart, err

}
