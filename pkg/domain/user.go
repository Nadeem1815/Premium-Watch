package domain

import "time"

type Users struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;not null"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	EmailId   string `json:"email_id" gorm:"uniqueIndex" binding:"required" validate:"required,email"`
	Phone     string `gorm:"uniqueIndex" json:"phone" validate:"required,phone"`
	Password  string `json:"password" binding:"required" validate:"required,min=8,max=64" gorm:"not null"`
	CreatedAt time.Time
}

type Address struct {
	ID        uint   `json:"id" gorm:"primaryKey;uinque;not null"`
	UsersID   string `json:"user_id"`
	Users     Users  `gorm:"foreignKey:UsersID"`
	HouseName string `json:"house_name" binding:"required"`
	Street    string `json:"street" binding:"required"`
	District  string `json:"district" binding:"required"`
	State     string `json:"state" binding:"required"`
	Landmark  string `json:"landmark" binding:"required"`
	PinCode   uint   `json:"pincode" binding:"requird"`
}

type UserInfo struct {
	ID        uint  `gorm:"primaryKey"`
	UsersID   uint  `json:"users_id"`
	Users     Users `gorm:"foreignKey:UsersID"`
	IsBlocked bool  `json:"is_blocked"`
	BlockedAt time.Time
	// BlockedBy        uint
	// Admin            Admin  `gorm:"foreignkey:BlockedBy"`
	// ResonforBlocking string `json:"reson_for_blocking"`
}

type Wallet struct {
	ID            uint    `json:"id" gorm:"primaryKey;not null"`
	UserID        uint    `json:"user_id" gorm:"not null"`
	Users         Users   `gorm:"foreignKey:UserID"`
	WalletBalance float64 `json:"total_amount" gorm:"not null"`
}
