package model

import (
	"time"
)

type NewCategory struct {
	CategoryName string `json:"category_name"`
}

type OutPutProduct struct {
	ID                uint
	ProductCategoryID uint
	Name              string
	Brand             string
	Colour            string
	Description       string
	Price             float64
	Stock             float64
	ProductImage      string
	SKU               string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CategoryName      string
}
