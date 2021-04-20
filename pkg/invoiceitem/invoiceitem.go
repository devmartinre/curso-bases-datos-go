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

// Storage interface
type Storage interface {
	Migrate() error
}

// Service of invoiceitem
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
