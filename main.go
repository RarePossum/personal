package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		
	})

	router.Run(":" + port)
}