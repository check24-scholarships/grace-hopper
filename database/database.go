package database

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func OpenDatabase() (*sql.DB, error){
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func CloseDatabase(db *sql.DB){
	db.Close()
}

func Search(db *sql.DB, name string) (Product){
	results, err := db.Query("SELECT * FROM products WHERE name LIKE '%?%'", name)
    if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &price, &image)
		if err != nil {
			log.Fatal(err)
		}
}
