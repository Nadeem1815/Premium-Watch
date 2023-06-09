package model

import "time"

type CreateCoupon struct {
	Code              string    `json:"code,omitempty"`
	MinOrderValue     float64   `json:"min_order_value"`
	DiscountPercent   float64   `json:"discount_percent"`
	DiscountMaxAmount float64   `json:"discount_max_amount"`
	ValidTill         time.Time `json:"valid_till"`
}

type UpdateCoupon struct {
	ID                int       `josn:"id"`
	Code              string    `gorm:"uinque" json:"code,omitempty"`
	DiscountPercent   float64   `json:"discount_percent"`
	DiscountMaxAmount float64   `json:"discount_max_amount"`
	ValidTill         time.Time `json:"valid_till"`
}
