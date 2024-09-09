package router

import (
	"training-go/go-session2/handler"
	"training-go/go-session2/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	userPublicEndpoint := r.Group("/users")
	userPublicEndpoint.GET("/", handler.GetAllUsers)
	userPublicEndpoint.GET("/:id", handler.GetUser)

	userPrivateEndpoint := r.Group("/users")
	userPrivateEndpoint.Use(middleware.AuthMiddleware())
	userPrivateEndpoint.POST("/", handler.CreateUser)
	userPrivateEndpoint.PUT("/:id", handler.UpdateUser)
	userPrivateEndpoint.DELETE("/:id", handler.DeleteUser)

}
