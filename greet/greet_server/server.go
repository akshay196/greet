package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/akshay196/grpc-demo/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet is invoked by %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	// lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName
	res := &pb.GreetResponse{
		Result: result,
	}
	return res, nil
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
