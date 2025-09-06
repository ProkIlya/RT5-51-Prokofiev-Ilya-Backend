// cmd/app/main.go
package main

import (
	"tesla-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func addSecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(addSecurityHeaders())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", handlers.IndexHandler)
	router.GET("/scenario/:id", handlers.ScenarioHandler)
	router.GET("/trip/:id", handlers.TripHandler)

	router.Run(":8080")
}
