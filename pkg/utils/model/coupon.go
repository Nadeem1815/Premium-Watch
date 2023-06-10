package model

import "time"

type CreatCoupon struct {
	Code              string    `json:"code,omitempty"`
	MinOrderValue     float64   `json:"min_order_value"`
	DiscountPercent   float64   `json:"discount_percent"`
	DiscountMaxAmount float64   `json:"discount_max_amount"`
	ValidTill         time.Time `json:"valid_till"`
}

type UpdatCoupon struct {
	ID                int       `josn:"id"`
	MinOrderValue     float64   `json:"min_order_value"`
	DiscountPercent   float64   `json:"discount_percent"`
	DiscountMaxAmount float64   `json:"discount_max_amount"`
	ValidTill         time.Time `json:"valid_till"`
}
