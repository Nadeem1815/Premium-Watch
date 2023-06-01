package interfaces

import (
	"context"

	"github.com/nadeem1815/premium-watch/pkg/domain"
)

type PaymentRepository interface {
	ViewPaymenDetails(ctx context.Context, orderID int) (domain.PaymentDetails, error)
	UpdatePaymentDetails(ctx context.Context, OrderID int, PaymentRef string) (domain.PaymentDetails, error)
}
