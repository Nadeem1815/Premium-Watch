package domain

import "time"

type Order struct {
	ID                uint      `gorm:"primaryKey;not null"`
	UserID            string    `json:"user_id"`
	Users             Users     `gorm:"foreignKey:UserID" json:"-"`
	OrderDate         time.Time `json:"order_date"`
	ShippingAddressID uint      `json:"shipping_address_id"`
	Address           Address   `gorm:"foreignKey:ShippingAddressID" json:"-"`
	OrderTotal        float64   `json:"order_total"`
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
	ID          uint `gorm:"primaryKey"`
	OrderStatus string
}
