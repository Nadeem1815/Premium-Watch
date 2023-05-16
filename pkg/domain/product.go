package domain

import "time"

type ProductCategory struct {
	ID           uint   `gorm:"primarykey,not null" json:"id"`
	CategoryName string `gorm:"not null,unique,index" json:"category_name"`
}

// type ProductBrand struct {
// 	ID                uint   `gorm:"primarykey,not null" json:"id"`
// 	Brand             string `gorm:"not null,index,unique" json:"brand"`
// 	Brand_Discription string `json:"brand_discription"`
// }

type Product struct {
	ID                uint            `gorm:"primarykey" json:"id"`
	ProductCategoryID uint            `json:"product_category_id"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID" json:"-"`
	Name              string          `json:"name" validate:"required"`
	Brand             string          `gorm:"not null" json:"brand" validate:"required"`
	Colour            string          `json:"colour"`
	Description       string          `json:"description"`
	Price             float64         `gorm:"not null" json:"price" validate:"required"`
	Stock             float64         `gorm:"not null" json:"stock" validate:"required"`
	ProductImage      string          `json:"product_image"`
	SKU               string          `gorm:"not null" json:"sku" validate:"required"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
