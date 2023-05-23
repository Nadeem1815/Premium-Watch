package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type CartRepository interface {
	AddToCart(ctx context.Context, userID string, productID int) (domain.CartItems, error)
	RemoveTOCart(ctx context.Context, userID string, productId int) error
	ViewCart(ctx context.Context, usertID string) (model.ViewCart, error)
}
