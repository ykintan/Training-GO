package router

import (
	"training-go/go-session1/handler"
	"training-go/go-session1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	// Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndpoint := r.Group("/api/v1")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		//privateEndpoint.POST("/post", handler.PostHandler)
		privateEndpoint.POST("/post", handler.PostHandler)
	}

}
