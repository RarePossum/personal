package main

import (
	"log"
	"net/http"
	"os"
	"io"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	gin.DisableConsoleColor()

    // Logging to a file.
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		
	})

	router.Run(":" + port)
}