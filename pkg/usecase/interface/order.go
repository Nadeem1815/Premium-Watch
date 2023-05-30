package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type OrderUseCase interface {
	BuyAll(ctx context.Context, body model.PlaceOrder, userID string) (domain.Order, error)
	CancelOrder(ctx context.Context, orderID int, UserID string) (domain.Order, error)
}
