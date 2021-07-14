package main

import (
	"context"
	"log"

	"github.com/akshay196/grpc-demo/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Default grpc run with SSL conneciton
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created clinet: %v", c)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Akshay",
			LastName:  "Gaikwad",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("erro while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
