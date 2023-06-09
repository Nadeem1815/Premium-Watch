package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type OrderUseCase interface {
	BuyAll(ctx context.Context, body model.PlaceOrder, userID string) (domain.Order, error)
	CancelOrder(ctx context.Context, orderID int, UserID string) (domain.Order, error)
	UpdateOrder(ctx context.Context, orderInfo model.UpdateOrder) (domain.Order, error)
	ViewAllOrder(ctx context.Context, UserID string) ([]domain.Order, error)
	ViewOrderID(ctx context.Context, userID string, orderID int) (domain.Order, error)
	ReturnReq(ctx context.Context, userID string, orderId model.RetrunRequest) (domain.Order, error)
	UserWallet(ctx context.Context, userID string) (domain.Wallet, error)
}
