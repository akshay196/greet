package main

import (
	"context"
	"fmt"
	"log"

	"github.com/akshay196/grpc-demo/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to connect to server: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doSum(c, 10, 3)
	// Error call
	doDivide(c, 10, 0)
	// Correct call
	doDivide(c, 10, 3)
}

func doSum(c calculatorpb.CalculatorServiceClient, num1, num2 int32) {
	req := &calculatorpb.SumRequest{
		Num1: num1,
		Num2: num2,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		fmt.Printf("Failed grpc response: %v", err)
		return
	}
	fmt.Printf("Result from grpc server (%d + %d): %d\n", num1, num2, res.GetSum())
}

func doDivide(c calculatorpb.CalculatorServiceClient, dividend, divisor int32) {
	res, err := c.Divide(context.Background(), &calculatorpb.DivisionRequest{
		DividendNumber: dividend,
		DivisorNumber:  divisor,
	})

	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// user's GRPC status error
			fmt.Println("Error code: ", resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("The divisor is probably 0")
			}
			return
		} else {
			// some big error
			log.Fatal(err)
			return
		}
	}
	fmt.Printf("Result from grpc server (%d / %d): %g\n", dividend, divisor, res.GetDivisionResult())
}
