package model

type DisplayCart struct {
	ID           int
	Brand        string
	Name         string
	Colour       string
	Quantity     uint
	ProductImage string
	Price        float64
	Total        float64
}

type ViewCart struct {
	CartItmes []DisplayCart `json:"cart_items,omitempty"`
	CouponID  int           `json:"coupon_id"`
	Discount  float64       `json:"discount"`
	SubTotal  float64       `json:"sub_total"`
	Total     float64       `json:"total,omitempty"`
}
