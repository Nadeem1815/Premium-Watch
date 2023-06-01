package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
)

type PaymentUseCase interface {
	CreateRazorPayment(ctx context.Context, userID string, orderID int) (domain.Order, string, error)
	UpatePaymentDetails(ctx context.Context, paymentVarifier model.PaymentVarification) error
}
