package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "training-go/go-session8-intro-grpc/proto/helloworld/v1"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &server{})
	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Message: fmt.Sprintf("Hello: %s", in.Name),
	}, nil
}
