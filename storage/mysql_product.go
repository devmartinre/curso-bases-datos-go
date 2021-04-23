package storage

import (
	"database/sql"
	"fmt"
)

const (
	mySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
)

// MySQLProduct used for work with mySQL - product
type MySQLProduct struct {
	db *sql.DB
}

// newMySQLProduct return a new pointer of MySQLProduct
func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

// Migrate implement the interface product.Storage
func (p *MySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de producto ejecutada correctamente")
	return nil
}

// Create implement the interface product.Storage
// func (p *MySQLProduct) Create(m *product.Model) error {
// 	stmt, err := p.db.Prepare(MySQLCreateProduct)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	err = stmt.QueryRow(
// 		m.Name,
// 		stringToNull(m.Observations),
// 		m.Price,
// 		m.CreatedAt,
// 	).Scan(&m.ID)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("se creo el producto correctamente")
// 	return nil
// }

// GetAll implement the interface product.Storage
// func (p *MySQLProduct) GetAll() (product.Models, error) {
// 	stmt, err := p.db.Prepare(MySQLGetAllProduct)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	ms := make(product.Models, 0)
// 	for rows.Next() {
// 		m, err := scanRowProduct(rows)
// 		if err != nil {
// 			return nil, err
// 		}
// 		ms = append(ms, m)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return ms, nil
// }

// GetByID implement the interface product.Storage
// func (p *MySQLProduct) GetByID(id uint) (*product.Model, error) {
// 	stmt, err := p.db.Prepare(MySQLGetProductByID)
// 	if err != nil {
// 		return &product.Model{}, err
// 	}
// 	defer stmt.Close()

// 	return scanRowProduct(stmt.QueryRow(id))
// }

// Update implement the interface product.Storage
// func (p *MySQLProduct) Update(m *product.Model) error {
// 	stmt, err := p.db.Prepare(MySQLUpdateProduct)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	res, err := stmt.Exec(
// 		m.Name,
// 		stringToNull(m.Observations),
// 		m.Price,
// 		timeToNull(m.UpdatedAt),
// 		m.ID,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, err := res.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAffected == 0 {
// 		return fmt.Errorf("no existe el producto con id: %d", m.ID)
// 	}

// 	fmt.Println("se actualizó el producto correctamente")
// 	return nil
// }

// Delete implement the interface product.Storage
// func (p *MySQLProduct) Delete(id uint) error {
// 	stmt, err := p.db.Prepare(MySQLDeleteProduct)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(id)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("se eliminó el producto correctamente")
// 	return nil
// }
