package domain

type PaymentDetails struct {
	ID              uint          `gorm:"primaryKey,index" json:"id"`
	OrderID         uint          `json:"oder_id,omitempty"`
	Order           Order         `gorm:"foreignKey:OrderID"`
	OrderTotal      float64       `json:"order_total"`
	PaymentMethodID uint          `json:"payment_method"`
	PaymentMethod   PaymentMethod `gorm:"foreignKey:PaymentMethodID"`
}

type PaymentMethod struct {
	ID            uint `gorm:"primaryKey"`
	PaymentMethod string
}