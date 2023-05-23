package model

type DisplayCart struct {
	ProductID    uint
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
	SubTotal  float64       `json:"sub_total"`
	Total     float64       `json:"total,omitempty"`
}
