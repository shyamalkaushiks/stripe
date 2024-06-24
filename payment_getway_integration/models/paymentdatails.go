package models

type Paymentdatails struct {
	Name   string  `form:"name" binding:"required"`
	Email  string  `form:"email" binding:"required"`
	Amount float64 `form:"amount" binding:"required"`
}

type TransactionDetails struct {
	CoustumerName string
	Amount        float64
	Status        string
}
