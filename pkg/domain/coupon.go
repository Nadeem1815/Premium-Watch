package domain

import "time"

type Coupon struct {
	ID                uint      `gorm:"primaryKey" json:"id,omitempty"`
	Code              string    `gorm:"uinque" json:"code,omitempty"`
	MinOrderValue     float64   `json:"min_order_value,omitempty"`
	DiscountPercent   float64   `json:"discount_percent"`
	DiscountMaxAmount float64   `json:"discount_max_amount"`
	ValidTill         time.Time `json:"valid_till"`
}
