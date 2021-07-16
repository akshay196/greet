package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/akshay196/grpc-demo/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	certFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("Failed while loading CA trust certificate: %v", sslErr)
		return
	}
	opts := grpc.WithTransportCredentials(creds)
	cc, err := grpc.Dial("localhost:50051", opts) // Default grpc run with SSL conneciton
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	// Greet
	// req := &greetpb.GreetRequest{
	// 	Greeting: &greetpb.Greeting{
	// 		FirstName: "Akshay",
	// 		LastName:  "Gaikwad",
	// 	},
	// }
	// res, err := c.Greet(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("erro while calling Greet RPC: %v", err)
	// }
	// log.Printf("Response from Greet: %v", res.Result)

	doGreetTimeout(c)
}

func doGreetTimeout(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Akshay",
			LastName:  "Gaikwad",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		sErr, ok := status.FromError(err)
		if ok {
			if sErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Time occurred!")
			}
		} else {
			fmt.Printf("Error in response: %v", sErr.Code())
		}
		return
	}
	log.Printf("Response from Greet: %v", res.Result)
}
