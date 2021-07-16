package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/akshay196/grpc-demo/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum is invoked by %v\n", req)
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	result := num1 + num2
	res := &calculatorpb.SumResponse{
		Sum: result,
	}
	return res, nil
}

func (*server) Divide(ctx context.Context, req *calculatorpb.DivisionRequest) (*calculatorpb.DivisionResponse, error) {
	fmt.Printf("Divide RPC is invoked: %v", req)
	divident := req.GetDividendNumber()
	divisor := req.GetDivisorNumber()

	if divisor == 0 {
		err := status.Errorf(codes.InvalidArgument, "Cannot divide by 0")
		return nil, err
	}
	division_result := divident / divisor
	res := &calculatorpb.DivisionResponse{
		DivisionResult: float64(division_result),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not announce listener: %v", err)
	}
	fmt.Println("Listening on :50051")

	s := grpc.NewServer()
	reflection.Register(s)
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start GRPC server: %v", err)
	}
}
