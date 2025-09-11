// cmd/app/main.go
package main

import (
	"tesla-app/internal/api"

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
	api.StartServer()
}
