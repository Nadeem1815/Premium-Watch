package model

type PaymentVarification struct {
	UserID     string
	OrderID    int
	PaymentRef string
	Total      float64
}
