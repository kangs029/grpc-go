package main

import (
	"context"
	"log"

	pb "github.com/kangs029/grpc-go/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	req := &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}
	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v\n", err)
	}
	log.Printf("Response from Sum: %v\n", res.Sum)
}
