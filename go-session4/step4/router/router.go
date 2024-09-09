package router

import (
	"training-go/go-session4/step4/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)
}
