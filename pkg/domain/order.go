package domain

import "time"

type Order struct {
	ID                uint           `gorm:"primaryKey;not null"`
	UserID            string         `json:"user_id"`
	Users             Users          `gorm:"foreignKey:UserID" json:"-"`
	OrderDate         time.Time      `json:"order_date"`
	PaymentMethodID   uint           `json:"payment_method_id"`
	PaymentMethod     PaymentMethod  `gorm:"foreignKey:PaymentMethodID" json:"-"`
	ShippingAddressID uint           `json:"shipping_address_id"`
	Address           Address        `gorm:"foreignKey:ShippingAddressID" json:"-"`
	OrderTotal        float64        `json:"order_total"`
	OrderStatusID     uint           `json:"order_status_id"`
	CouponID          uint           `json:"coupon_id"`
	OrderStatus       OrderStatus    `gorm:"foreignKey:OrderStatusID" json:"-"`
	DeliveryStatusID  uint           `json:"delivery_status_id"`
	DeliveryStatus    DeliveryStatus `gorm:"foreignKey:DeliveryStatusID"`
	DeliveryUpdatedAt time.Time      `json:"delivery_time"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"-"`
	OrderID   uint    `json:"order_id"`
	Order     Order   `gorm:"foriegnKey:OrderID" json:"-"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderStatus struct {
	ID           uint   `gorm:"primaryKey"`
	Order_Status string `json:"order_status"`
}

type DeliveryStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}
