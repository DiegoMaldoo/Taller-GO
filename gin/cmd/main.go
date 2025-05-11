package main

import (
	"fmt"
	"gin/cmd/middleware"
	"gin/cmd/routes"
	"gin/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// middleware global
	r.Use(middleware.LoggerMiddleware())

	// creamos una sola instancia de UserService
	userService := services.NewUserService()

	// ruta de prueba
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "ok",
		})
	})

	routes.SetupUserRoutes(r, userService)

	fmt.Println("Listening at port 3000")
	r.Run(":3000")
}
