package main

import (
	"context"
	"log"

	"training-go/go-session6-db-pgx-crud/handler"
	postgrespgx "training-go/go-session6-db-pgx-crud/repository/postgres_pgx"
	"training-go/go-session6-db-pgx-crud/router"
	"training-go/go-session6-db-pgx-crud/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	//setup service
	dsn := "postgresql://postgres:P4ssw0rd@localhost:5432/training_golang"
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := postgrespgx.NewUserRepository(pool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Server started at localhost:8080")
	r.Run("localhost:8080")
}
