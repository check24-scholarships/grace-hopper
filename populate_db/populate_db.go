package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func main() {
	jsonPath := "populate_db/dump.json"
	rawData, err := ioutil.ReadFile(jsonPath)

	path, _ := os.Getwd()
	fmt.Println(path)

	if err != nil {
		log.Fatal(err)
		return
	}

	var table []Product
	err = json.Unmarshal([]byte(rawData), &table)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, product := range table {
		fmt.Println(product.Name)
		fmt.Println(product.Price)
		fmt.Println(product.Image)
		fmt.Println("---")
	}
}
