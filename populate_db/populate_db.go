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

	var table []Product
	table, err := ReadTableFromJson(jsonPath)

	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := database.OpenDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer database.CloseDatabase(db)

	for _, product := range table {
		err = database.InsertProduct(db, database.Product(product))

		if err != nil {
			log.Println("Warning: Unable to insert product into DB")
			log.Println(err)
		}
	}
}
