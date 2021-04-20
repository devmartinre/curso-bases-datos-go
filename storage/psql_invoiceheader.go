package storage

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
)

// PsqlInvoiceHeader used to work with postgres - product
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal("An error happened")
		}
	}(stmt)

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migration of invoice header executed successfully")
	return nil
}
