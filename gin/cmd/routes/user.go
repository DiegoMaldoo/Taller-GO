// user_routes.go
package routes

import (
	"gin/cmd/controllers"
	"gin/cmd/middleware"
	"gin/cmd/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userService *services.UserService) {
	// Creamos el controller
	uc := controllers.NewUserController(userService)

	// Grupo de rutas /users con middleware de API key
	grp := r.Group("/users")
	grp.Use(middleware.APIKeyAuthMiddleware())

	grp.GET("", uc.GetUsers)
	grp.POST("", uc.CreateUser)
	grp.PUT("/:id", uc.UpdateUser)
	grp.DELETE("/:id", uc.DeleteUser)
}
