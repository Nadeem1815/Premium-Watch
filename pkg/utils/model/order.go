package model

type PlaceOrder struct {
	// ProductID         uint `json:"product_id,omitempty"`
	PaymentMethodID   uint `json:"payment_method_id,omitempty"`
	ShippingAddressID int  `json:"shipping_address_id"`
}

type OrderBuyItem struct {
	ProductID         uint `json:"product_id"`
	ShippingAddressID uint `json:"shipping_address_id"`
}

type UpdateOrder struct {
	OrderID          uint `json:"order_id"`
	OrderStatusID    uint `json:"status_id"`
	DeliveryStatusId uint `json:"delivery_status_id"`
}
