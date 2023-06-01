package usecase

import (
	"context"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/razorpay/razorpay-go"
)

const (
	razorpayID     = "rzp_test_kbr5PARM3lK7xo"
	razorpaysecret = "l0bzJw57T8IBHEXBv5re0DIm"
)

type paymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
	orderRepo   interfaces.OrderRepoitory
}

func NewPaymentUseCase(orderRepo interfaces.OrderRepoitory, paymentRepo interfaces.PaymentRepository) services.PaymentUseCase {
	return &paymentUseCase{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
	}
}

func (cr *paymentUseCase) CreateRazorPayment(ctx context.Context, userID string, orderID int) (domain.Order, string, error) {
	// check if payment already is paid  no need proceed payment. if not paid yet proceed with transaction
	paymentDetails, err := cr.paymentRepo.ViewPaymenDetails(ctx, orderID)
	if err != nil {
		return domain.Order{}, "", err

	}
	if paymentDetails.PaymentStatusID == 2 {
		return domain.Order{}, "", fmt.Errorf("payment already completed")
	}
	// fetch order details from databse
	order, err := cr.orderRepo.ViewOrderID(ctx, userID, orderID)
	if err != nil {
		return domain.Order{}, "", err

	}
	if order.ID == 0 {
		return domain.Order{}, "", fmt.Errorf("no orders")

	}
	client := razorpay.NewClient(razorpayID, razorpaysecret)

	data := map[string]interface{}{
		"amount":   order.OrderTotal * 100,
		"currency": "INR",
		"receipt":  "test_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Order{}, "", err

	}
	value := body["id"]
	razorpayID := value.(string)
	return order, razorpayID, err

}

func (cr *paymentUseCase) UpatePaymentDetails(ctx context.Context, paymentVarifier model.PaymentVarification) error {
	//  fetch payment details
	paymentDetails, err := cr.paymentRepo.ViewPaymenDetails(ctx, paymentVarifier.OrderID)
	if err != nil {
		return err

	}
	if paymentDetails.ID == 0 {
		return fmt.Errorf("no order found")

	}
	if paymentDetails.OrderTotal != paymentVarifier.Total {
		return fmt.Errorf("payment and order amount does not match")

	}
	updatePayment, err := cr.paymentRepo.UpdatePaymentDetails(ctx, paymentVarifier.OrderID, paymentDetails.PaymentRef)
	if err != nil {
		return err

	}
	if updatePayment.ID == 0 {
		return fmt.Errorf("failed update payment details")
	}
	return nil
}
