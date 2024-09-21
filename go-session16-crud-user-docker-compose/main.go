package main

import (
	"context"
	"log"
	"net"

	"go-session16-crud-user-compose/entity"
	grpcHandler "go-session16-crud-user-compose/handler/grpc"
	"go-session16-crud-user-compose/middleware"
	pb "go-session16-crud-user-compose/proto"
	postgresgormraw "go-session16-crud-user-compose/repository/postgres_gorm_raw"
	"go-session16-crud-user-compose/service"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	//setup service
	dsn := "postgresql://postgres:P4ssw0rd@pg-db:5432/training_golang"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	if err := gormDB.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalln("Failed to migrate database : ", err)
	}

	userRepo := postgresgormraw.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := grpcHandler.NewUserHandler(userService)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor((middleware.UnaryAuthInterceptor())))
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		log.Println("Server is running on port :50051")
		grpcServer.Serve(lis)
	}()

	//run grpc gateway
	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server :", err)
	}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway :", err)
	}

	//run gin server
	ginserver := gin.Default()
	ginserver.Group("/v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))

	log.Println("Server is running on port :8080")

	ginserver.Run("0.0.0.0:8080")
}
