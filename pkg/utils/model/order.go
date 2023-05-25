package model

type PlaceOrder struct {
	ProductID         uint `json:"product_id,omitempty"`
	ShippingAddressID int  `json:"shipping_address_id"`
}
