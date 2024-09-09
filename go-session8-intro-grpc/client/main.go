package main

import (
	"context"
	"log"
	pb "training-go/go-session8-intro-grpc/proto/helloworld/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	greeterClient := pb.NewGreeterServiceClient(conn)

	res, err := greeterClient.SayHello(context.Background(), &pb.SayHelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}
	log.Println("SayHello response:", res.Message)

}
