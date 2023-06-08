package model

type AdminLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,email"`
}

type AdminDataOutput struct {
	ID       string `gorm:"primaryKey,index" json:"id"`
	UserName string `gorm:"uniqueIndex" json:"user_name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
}

type AdminDashBoard struct {
	CompletedOrders int     `json:"completed_orders,omitempty"`
	PendingOrders   int     `json:"pending_orders,omitempty"`
	CancelledOrders int     `json:"cancelled_orders,omitempty"`
	TotalOrders     int     `json:"total_orders,omitempty"`
	TotalOrderItems int     `json:"total_order_items,omitempty"`
	OrderValue      float64 `json:"order_value,omitempty"`
	CreditedAmount  float64 `json:"credited_amount,omitempty"`
	PendingAmount   float64 `json:"pending_amount,omitempty"`

	TotalUsers   int `json:"total_users,omitempty"`
	OrderedUsers int `json:"ordered_users,omitempty"`
}
