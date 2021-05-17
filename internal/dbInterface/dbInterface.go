package dbInterface

import (
	"restserver/internal/structs/companystruct"
	"restserver/internal/structs/productstruct"
)

type Database interface {
	DatabaseProducts
	DatabaseCompany
}
type DatabaseProducts interface {
	GetAllProducts() []productstruct.Product
	AddProduct(product productstruct.Product) bool
	UpdateProduct(product productstruct.Product) bool
	GetProduct(id int) (productstruct.Product, bool)
	DeleteProduct(id int) bool
}

type DatabaseCompany interface {
	GetAllCompanys() []companystruct.Company
	AddCompany(company companystruct.Company) bool
	DeleteCompany(id int) bool
	GetAllProductsInCompany(companyName string) ([]productstruct.Product, bool)
}
