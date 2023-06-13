package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type CartUseCase interface {
	AddToCart(ctx context.Context, userID string, productID int) (domain.CartItems, error)
	RemoveTOCart(ctx context.Context, userID string, productId int) error
	ViewCart(ctx context.Context, userID string) (model.ViewCart, error)

	AddCouponToCart(ctx context.Context, userID string, couponID int) (model.ViewCart, error)
}
