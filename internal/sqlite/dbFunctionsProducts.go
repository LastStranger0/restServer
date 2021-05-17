package sqlite

import (
	"database/sql"
	"log"
	"restserver/internal/structs/productstruct"
	"strconv"
)

var PostgreSQL postgresSQL

func init() {
	dbConnect, err := sql.Open("sqlite3", "MyDb.db")
	if err != nil {
		panic(err)
	}
	PostgreSQL = postgresSQL{
		dbConnection: dbConnect,
	}
}

const createTableProductsExec = "CREATE TABLE products " +
	"(id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, " +
	"model varchar(30) not null, " +
	"company varchar(30) not null, " +
	"price int not null)"

func (p postgresSQL) CreateTableProducts() {
	_, err := p.dbConnection.Exec(createTableProductsExec)
	if err != nil {
		log.Println("CreateTableProducts", err)
	}
}

const getAllProductsQuery = `SELECT * FROM products`

func (p postgresSQL) GetAllProducts() []productstruct.Product {
	rows, err := p.dbConnection.Query(getAllProductsQuery)
	if err != nil {
		log.Fatal("GetAllProducts: ", err)
		return []productstruct.Product{}
	}
	defer rows.Close()

	products := make([]productstruct.Product, 0, 10)
	for rows.Next() {
		product := productstruct.Product{}
		err = rows.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
		if err != nil {
			log.Println("GetAllProducts: parse err", err)
			continue
		}
		products = append(products, product)
	}

	return products
}

const addProductExec = `insert into Products (model, company, price) values (?,?,?)`

func (p postgresSQL) AddProduct(product productstruct.Product) bool {
	successful := true
	_, err := p.dbConnection.Exec(addProductExec, product.Model, product.Company, strconv.Itoa(product.Price))
	if err != nil {
		log.Println("AddProduct ", err)
		successful = false
	}

	return successful
}

const updateProductExec = `update Products set model = ?, company = ?, price = ? where id = ?`

func (p postgresSQL) UpdateProduct(product productstruct.Product) bool {
	successful := true
	_, err := p.dbConnection.Exec(updateProductExec, product.Model, product.Company,
		strconv.Itoa(product.Price), strconv.Itoa(product.Id))
	if err != nil {
		log.Println("UpdateProduct ", err)
		successful = false
	}

	return successful
}

const getProductQueryRow = `select * from Products where id = ?`

func (p postgresSQL) GetProduct(id int) (productstruct.Product, bool) {
	successful := true
	row := p.dbConnection.QueryRow(getProductQueryRow, strconv.Itoa(id))
	product := productstruct.Product{}
	err := row.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
	if err != nil {
		log.Println("GetProduct ", err)
		successful = false
	}

	return product, successful
}

const deleteProductExec = `delete from Products where id = ?`

func (p postgresSQL) DeleteProduct(id int) bool {
	successful := true
	_, err := p.dbConnection.Exec(deleteProductExec, strconv.Itoa(id))
	if err != nil {
		log.Println("DeleteProduct ", err)
		successful = false
	}

	return successful
}

const getAllProductsInCompanyQuery = `select * from Products where company = ?`

func (p postgresSQL) GetAllProductsInCompany(companyName string) ([]productstruct.Product, bool) {
	rows, err := p.dbConnection.Query(getAllProductsInCompanyQuery, companyName)
	if err != nil {
		log.Fatal("GetAllProducts: ", err)
		return []productstruct.Product{}, false
	}
	defer rows.Close()

	products := make([]productstruct.Product, 0, 10)
	for rows.Next() {
		product := productstruct.Product{}
		err = rows.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
		if err != nil {
			log.Println("GetAllProductsInCompany: parse err", err)
			continue
		}
		products = append(products, product)
	}

	return products, true
}
