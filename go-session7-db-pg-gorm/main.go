package main

import (
	"log"

	"training-go/go-session7-db-pg-gorm/handler"
	postgresgormraw "training-go/go-session7-db-pg-gorm/repository/postgres_gorm_raw"
	"training-go/go-session7-db-pg-gorm/router"
	"training-go/go-session7-db-pg-gorm/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	//setup service
	dsn := "postgresql://postgres:P4ssw0rd@localhost:5432/training_golang"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := postgresgormraw.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Server started at localhost:8080")
	r.Run("localhost:8080")
}
