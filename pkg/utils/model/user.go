package model

type UsarDataInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	EmailId  string `json:"email_id" binding:"required"  validate:"required,email"`
	Phone    string `json:"phone" validate:"required,phone"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=64"`
}

type UserDataOutput struct {
	ID      uint   `json:"user_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	EmailId string `json:"email_id"`
	Phone   string `json:"phone"`
}

type UserLoginEmail struct {
	EmailId  string `json:"email_id" binding:"required"  validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=64" gorm:"not null"`
}

type UserLoginVarifier struct {
	ID        uint   `json:"user_id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	EmailId   string `json:"email_id"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	IsBlocked bool   `json:"is_blocked"`
}

type BlockUser struct {
	UsarId int    `json:"user_id"`
	Reason string `json:"reason"`
}

type AddressInput struct {
	HouseName string `json:"house_name"`
	Street    string `json:"street"`
	District  string `json:"district"`
	State     string `json:"state"`
	Landmark  string `json:"landmark"`
	PinCode   uint   `json:"pincode"`
}
