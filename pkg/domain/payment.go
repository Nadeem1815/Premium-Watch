package domain

import "time"

type PaymentStatus struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Status string `json:"status"`
}

type PaymentMethod struct {
	ID            uint `gorm:"primaryKey"`
	PaymentMethod string
}

type PaymentDetails struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	OrderID         uint          `json:"oder_id,omitempty"`
	Order           Order         `gorm:"foreignKey:OrderID"`
	OrderTotal      float64       `json:"order_total"`
	PaymentMethodID uint          `json:"payment_method"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
	PaymentStatusID uint          `json:"payment_status_id"`
	PaymentStatus   PaymentStatus `gorm:"foreignKey:PaymentStatusID"`
	PaymentRef      string        `gorm:"uinque"`
	UpdatedAt       time.Time
}
