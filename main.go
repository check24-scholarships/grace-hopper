package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		datetime := time.Now().Format("01-02-2006 15:04:05")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"time": datetime,
		})
	})
	router.Run(":8000")

}
