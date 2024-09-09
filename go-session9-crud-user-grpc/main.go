package main

import (
	"log"
	"net"

	grpcHandler "training-go/go-session9-crud-user-grpc/handler/grpc"
	pb "training-go/go-session9-crud-user-grpc/proto"
	postgresgormraw "training-go/go-session9-crud-user-grpc/repository/postgres_gorm_raw"
	"training-go/go-session9-crud-user-grpc/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	//setup service
	dsn := "postgresql://postgres:P4ssw0rd@localhost:5432/training_golang"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := postgresgormraw.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := grpcHandler.NewUserHandler(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server is running on port :50051")
	grpcServer.Serve(lis)

}
