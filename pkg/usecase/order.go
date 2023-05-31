package usecase

import (
	"context"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type OrderUseCase struct {
	orderRepo   interfaces.OrderRepoitory
	userRepo    interfaces.UserRepository
	productRepo interfaces.ProductRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepoitory, userRepo interfaces.UserRepository, productRepo interfaces.ProductRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (cr *OrderUseCase) BuyAll(ctx context.Context, body model.PlaceOrder, UserID string) (domain.Order, error) {
	// check if user has added address. if not, retrun error
	address, err := cr.userRepo.ViewAddress(ctx, UserID)
	if err != nil {
		return domain.Order{}, err
	}
	if address.ID == 0 {
		return domain.Order{}, fmt.Errorf("cannot placed order without adding address")
	}
	orders, err := cr.orderRepo.BuyAll(ctx, body, UserID)
	return orders, err
}

func (cr *OrderUseCase) CancelOrder(ctx context.Context, orderID int, UserID string) (domain.Order, error) {
	cancelOrder, err := cr.orderRepo.CancelOrder(ctx, orderID, UserID)
	return cancelOrder, err
}

func (cr *OrderUseCase) UpdateOrder(ctx context.Context, orderInfo model.UpdateOrder) (domain.Order, error) {
	updatedOrder, err := cr.orderRepo.UpdateOrder(ctx, orderInfo)
	return updatedOrder, err
}
