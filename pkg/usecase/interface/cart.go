package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
)

type CartUseCase interface {
	AddToCart(ctx context.Context, userID string, productID int) (domain.CartItems, error)
}
