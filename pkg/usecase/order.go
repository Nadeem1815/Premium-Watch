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

func (cr *OrderUseCase) ViewAllOrder(ctx context.Context, UserID string) ([]domain.Order, error) {
	viewOrder, err := cr.orderRepo.ViewAllOrder(ctx, UserID)
	return viewOrder, err
}

func (cr *OrderUseCase) ViewOrderID(ctx context.Context, userID string, orderID int) (domain.Order, error) {
	viewOrderID, err := cr.orderRepo.ViewOrderID(ctx, userID, orderID)
	return viewOrderID, err
}

func (cr *OrderUseCase) ReturnReq(ctx context.Context, userID string, retrurnReqst model.RetrunRequest) (domain.Order, error) {
	// check if order is eligible to be returned

	orderDetails, err := cr.orderRepo.ViewOrderID(ctx, userID, retrurnReqst.OrderID)
	if err != nil {
		return domain.Order{}, err

	}
	if orderDetails.ID == 0 {
		return domain.Order{}, fmt.Errorf("no orders")

	}
	if orderDetails.DeliveryUpdatedAt.Sub(time.Now()) > time.Hour*24*14 {
		return domain.Order{}, fmt.Errorf("failed retrun your order more then 15 days ")
	}
	if orderDetails.OrderStatusID != 1 || orderDetails.DeliveryStatusID != 2 {
		return domain.Order{}, fmt.Errorf("cannot return order status %v and delivery status %v", orderDetails.OrderStatusID, orderDetails.DeliveryStatusID)

	}
	orderId, err := cr.orderRepo.ReturnReq(ctx, retrurnReqst)
	if err != nil {
		return domain.Order{}, fmt.Errorf("request failed")

	}
	return orderId, nil

}
