package model

type Invoice struct {
	ID            int     `db:"id" json:"id"`
	CompanyID     int     `db:"company_id" json:"company_id"`
	ClientID      int     `db:"client_id" json:"client_id"`
	IssueDate     string  `db:"issue_date" json:"issue_date"`
	PaymentAmount uint64  `db:"payment_amount" json:"payment_amount"`
	FeeAmount     uint64  `db:"fee" json:"fee_amount"`
	FeeRate       float64 `db:"fee_rate" json:"fee_rate"`
	TaxAmount     uint64  `db:"tax" json:"tax_amount"`
	TaxRate       float64 `db:"tax_rate" json:"tax_rate"`
	TotalAmount   uint64  `db:"total_amount" json:"total_amount"`
	DueDate       string  `db:"due_date" json:"due_date"`
	Status        string  `db:"status" json:"status"`
}

type Status string

const (
	InvoiceStatusPending   Status = "pending"
	InvoiceStatusProcessed Status = "processed"
	InvoiceStatusPaid      Status = "paid"
	InvoiceStatusFailed    Status = "failed"
)
