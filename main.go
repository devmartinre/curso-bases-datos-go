package main

import (
	"fmt"
	"log"

	"github.com/whiteagleinc-meli/curso-bases-datos-go/pkg/product"
	"github.com/whiteagleinc-meli/curso-bases-datos-go/storage"
)

func main() {
	storage.NewPostgresDB()

	//storageProduct := storage.NewPsqlProduct(storage.Pool())
	//serviceProduct := product.NewService(storageProduct)
	//
	//if err := serviceProduct.Migrate(); err != nil {
	//	log.Fatalf("product.Migrate: %v", err)
	//}

	//storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	//serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	//
	//if err := serviceInvoiceHeader.Migrate(); err != nil {
	//	log.Fatalf("serviceInvoiceHeader.Migrate: %v", err)
	//}

	//storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	//serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	//
	//if err := serviceInvoiceItem.Migrate(); err != nil {
	//	log.Fatalf("serviceInvoiceItem.Migrate: %v", err)
	//}

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m := product.Model{
	// 	Name:         "Curso de BD con Go",
	// 	Price:        70,
	// 	Observations: "On Fire!",
	// }

	// if err := serviceProduct.Create(&m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// fmt.Printf("%+v\n", m)

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// ms, err := serviceProduct.GetAll()

	// if err != nil {
	// 	log.Fatalf("product.GetAll: %v", err)
	// }

	// fmt.Println(ms)

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(2)

	if err != nil {
		log.Fatalf("product.GetByID: %v", err)
	}

	fmt.Println(m)
}
