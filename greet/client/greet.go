package main

import (
	"context"
	"log"

	pb "github.com/kangs029/grpc-go/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	// 1. Prepare the request
	req := &pb.GreetRequest{
		FirstName: "Kang",
	}

	// 2. Call Greet
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet: %v", err)
	}

	// 3. Print the response
	log.Printf("Response from Greet: %v", res.Result)
}
