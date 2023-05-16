package model

type AdminLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,email"`
}

type AdminDataOutput struct {
	ID       uint   `gorm:"primaryKey,index" json:"id"`
	UserName string `gorm:"uniqueIndex" json:"user_name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
}
