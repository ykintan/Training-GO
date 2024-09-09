package router

import (
	"training-go/go-session5-validator/handler"
	"training-go/go-session5-validator/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler) {
	usersPublicEndpoint := r.Group("/users")

	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)

	userPrivateEndpoint := r.Group("/users")
	userPrivateEndpoint.Use(middleware.AuthMiddleware())

	userPrivateEndpoint.POST("", userHandler.CreatedUser)
	userPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
	userPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)
}
