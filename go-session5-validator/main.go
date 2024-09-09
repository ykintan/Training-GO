package main

import (
	"log"

	"training-go/go-session5-validator/entity"
	"training-go/go-session5-validator/handler"
	"training-go/go-session5-validator/repository/slice"
	"training-go/go-session5-validator/router"
	"training-go/go-session5-validator/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	//setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Server started at localhost:8080")
	r.Run("localhost:8080")
}
