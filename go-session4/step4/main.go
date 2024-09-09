package main

import (
	"log"
	"training-go/go-session4/step4/entity"
	"training-go/go-session4/step4/handler"
	"training-go/go-session4/step4/repository/slice"
	"training-go/go-session4/step4/router"
	"training-go/go-session4/step4/service"

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

	//setup router
	router.SetupRouter(r, userHandler)

	log.Println("Server started at localhost:8080")
	r.Run("localhost:8080")
}
