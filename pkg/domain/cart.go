package domain

type Cart struct {
	ID       uint    `gorm:"primaryKey,index" json:"id"`
	UserID   string  `json:"user_id"`
	Users    Users   `gorm:"foreignKey:UserID" json:"-"`
	CouponID uint     `json:"coupon_id"`
	SubTotal float64 `json:"sub_total"`
	Discount float64 `json:"discount"`
	Total    float64 `json:"total"`
}

type CartItems struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `json:"cart_id"`
	Cart      Cart    `gorm:"foreignKey:CartID" json:"-"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"-"`
	Quantity  uint    `json:"quantity"`
	ItemTotal uint
}
