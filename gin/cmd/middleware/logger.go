package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// startTime := time.Now()
		method := c.Request.Method
		path := c.Request.Method
		clientIP := c.ClientIP()
		log.Printf("Request: %s %s from %s", method, path, clientIP)
		c.Next()
	}
}

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		// Reemplaza "my-secret-api-key" por la clave deseada
		if apiKey != "my-secret-api-key" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}
		c.Next()
	}
}
