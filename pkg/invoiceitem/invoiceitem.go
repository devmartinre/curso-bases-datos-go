package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID            uint
	InvoiceHeader uint
	ProductID     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
