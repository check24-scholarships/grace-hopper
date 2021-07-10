package main

import (
	"encoding/json"
	"grace-hopper/database"
	"io/ioutil"
	"log"
)

type Product = database.Product

func ReadTableFromJson(fileName string) ([]Product, error) {
	rawData, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var table []Product
	err = json.Unmarshal([]byte(rawData), &table)

	if err != nil {
		return nil, err
	}

	return table, nil
}

func main() {
	jsonPath := "populate_db/dump.json"

	log.Println("Reading product data from JSON")
	var table []Product
	table, err := ReadTableFromJson(jsonPath)

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Establishing DB connection")
	db, err := database.OpenDatabase()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer database.CloseDatabase(db)

	log.Printf("Inserting %d products\n", len(table))
	log.Println()

	for idx, product := range table {
		log.Printf("[%d/%d]", idx+1, len(table))
		err = database.InsertProduct(db, database.Product(product))

		if err != nil {
			log.Println("Warning: Unable to insert product into DB")
			log.Println(err)
		}
	}

	log.Println("Inserted the products")
}
