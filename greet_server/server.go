package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/akshay196/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
