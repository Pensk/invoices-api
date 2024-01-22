package model

import "time"

type Invoice struct {
	ID            int
	CompanyID     int
	ClientID      int
	IssueDate     time.Time
	PaymentAmount uint64
	FeeAmount     uint64
	FeeRate       float64
	TaxAmount     uint64
	TaxRate       float64
	TotalAmount   uint64
	DueDate       time.Time
	Status        string
}

type status string

const (
	Unpaid   status = "unpaid"
	Paid     status = "paid"
	Canceled status = "canceled"
	Error    status = "error"
)
