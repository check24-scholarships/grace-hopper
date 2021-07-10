package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func CloseDatabase(db *sql.DB) {
	db.Close()
}

func InsertProduct(db *sql.DB, product Product) error {
	query := "INSERT INTO products (`name`, `price`, `image`) VALUES (?, ?, ?)"
	_, err := db.Query(query, product.Name, product.Price, product.Image)
	return err
}

func Search(db *sql.DB, name string) []Product {
	rows, err := db.Query("SELECT * FROM products WHERE name LIKE '%?%'", name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	var currentProduct Product

	for rows.Next() {
		err := rows.Scan(&currentProduct.Name, &currentProduct.Price, &currentProduct.Image)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, currentProduct)
	}
	return products
}
