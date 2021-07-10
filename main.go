package main

import (
	"log"
	"net/http"
	"time"

	"grace-hopper/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	db, err := database.OpenDatabase()

	if err != nil {
		log.Println("Unable to establish a connection to the DB")
		log.Fatal(err)
	}

	defer database.CloseDatabase(db)

	router.GET("/", func(c *gin.Context) {
		datetime := time.Now().Format("01-02-2006 15:04:05")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"time": datetime,
		})
	})
	router.GET("/search", func(c *gin.Context) {
		searchQuery, _ := c.GetQuery("q")
		products := database.Search(db, searchQuery)

		c.HTML(http.StatusOK, "search.html", gin.H{
			"products": products,
		})
	})

	router.Run(":8000")

}
