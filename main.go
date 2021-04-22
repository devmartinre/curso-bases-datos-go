package main

import (
	"github.com/whiteagleinc-meli/curso-bases-datos-go/storage"
)

func main() {
	//storage.New(storage.Postgres)

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

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m, err := serviceProduct.GetByID(3)

	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("There is not a product with this ID")
	// case err != nil:
	// 	log.Fatalf("product.GetByID: %v", err)
	// default:
	// 	fmt.Println(m)
	// }

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m := &product.Model{
	// 	ID:           5,
	// 	Name:         "Curso de Python",
	// 	Observations: "This is Python Course",
	// 	Price:        120,
	// }

	// err := serviceProduct.Update(m)

	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// err := serviceProduct.Delete(3)

	// if err != nil {
	// 	log.Fatalf("product.Delete: %v", err)
	// }

	// storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	// storageInvoice := storage.NewPsqlInvoice(
	// 	storage.Pool(),
	// 	storageHeader,
	// 	storageItems,
	// )

	// m := &invoice.Model{
	// 	Header: &invoiceheader.Model{
	// 		Client: "Martin Rangel",
	// 	},
	// 	Items: invoiceitem.Models{
	// 		&invoiceitem.Model{ProductID: 4},
	// 		&invoiceitem.Model{ProductID: 5},
	// 	},
	// }

	// serviceInvoice := invoice.NewService(storageInvoice)
	// if err := serviceInvoice.Create(m); err != nil {
	// 	log.Fatalf("invoice.Create: %v", err)
	// }

	// serviceInvoice = invoice.NewService(storageInvoice)
	// if err := serviceInvoice.Create(m); err != nil {
	// 	log.Fatalf("invoice.Create: %v", err)
	// }

	storage.New(storage.MySQL)
}
