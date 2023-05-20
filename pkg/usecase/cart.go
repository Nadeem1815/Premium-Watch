package usecase

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
)

type CartUseCase struct {
	cartRepo interfaces.CartRepository
}

func NewCartUseCase(repo interfaces.CartRepository) services.CartUseCase {
	return &CartUseCase{
		cartRepo: repo,
	}
}

func (cr *CartUseCase) AddToCart(ctx context.Context, userID string, productID int) (domain.CartItems, error) {
	cartItem, err := cr.cartRepo.AddToCart(ctx, userID, productID)
	return cartItem, err
}
