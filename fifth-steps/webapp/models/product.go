package models

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func ListProduct(db *sql.DB) []Product {
	query := "SELECT id, name, description, price, quantity FROM products;"
	result, err := db.Query(query)

	if err != nil {
		panic(err)
	}
	var products []Product

	for result.Next() {
		var product Product
		err := result.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)

		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products
}

func GetProduct(db *sql.DB, id int) Product {
	query := "SELECT id, name, description, price, quantity FROM products WHERE id = $1;"
	sqlRow := db.QueryRow(query, id)
	if sqlRow.Err() != nil {
		panic(sqlRow.Err())
	}
	var product Product
	sqlRow.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
	return product
}

func (product *Product) CreateProduct(db *sql.DB) (sql.Result, error) {
	query := `INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4);`
	prdStr := fmt.Sprint(product)
	println("Product to be transmitted: ", prdStr)
	return db.Exec(query, product.Name, product.Description, product.Price, product.Quantity)
}

func (product *Product) DeleteProduct(db *sql.DB) (sql.Result, error) {
	query := `DELETE FROM products WHERE id = $1;`
	return db.Exec(query, product.ID)
}

func (product *Product) UpdateProduct(db *sql.DB) (sql.Result, error) {
	query := `UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5;`
	return db.Exec(query, product.Name, product.Description, product.Price, product.Quantity, product.ID)
}
