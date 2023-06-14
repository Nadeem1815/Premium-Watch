package repository

import (
	"context"
	"fmt"

	"github.com/nadeem1815/premium-watch/pkg/domain"
	interfaces "github.com/nadeem1815/premium-watch/pkg/repository/interface"
	"gorm.io/gorm"
)

type paymentDataBase struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentDataBase{DB}
}

func (cr *paymentDataBase) ViewPaymenDetails(ctx context.Context, orderID int) (domain.PaymentDetails, error) {
	var paymentDtls domain.PaymentDetails

	err := cr.DB.Raw("SELECT * FROM payment_details WHERE order_id = $1 LIMIT 1", orderID).Scan(&paymentDtls).Error
	if err != nil {
		return domain.PaymentDetails{}, err
	}
	return paymentDtls, nil

}

func (cr *paymentDataBase) UpdatePaymentDetails(ctx context.Context, OrderID int, PaymentRef string) (domain.PaymentDetails, error) {
	var updatePayment domain.PaymentDetails
	fmt.Println("repo ref 3", PaymentRef)
	updatePaymentQuery := `UPDATE payment_details 
								SET payment_method_id = 2, 
									payment_status_id = 2, 										payment_ref = $1, 
									updated_at = NOW() 
								WHERE order_id = $2 
								RETURNING *;`
	err := cr.DB.Raw(updatePaymentQuery, PaymentRef, OrderID).Scan(&updatePayment).Error
	if err != nil {
		return domain.PaymentDetails{}, err

	}
	return updatePayment, nil
}
