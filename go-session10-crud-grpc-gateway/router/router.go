package router

import (
	ginHandler "training-go/go-session9-crud-user-grpc/handler/gin"
	"training-go/go-session9-crud-user-grpc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler ginHandler.IUserHandler) {
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
