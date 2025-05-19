package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/kangs029/grpc-go/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	fmt.Println("doAverage was invoked")

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average RPC: %v", err)
	}

	numbers := []int32{1, 2, 3, 4, 5}
	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)
		stream.Send(&pb.AverageRequest{
			Number: number,
		})
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average RPC: %v", err)
	}

	fmt.Printf("Average: %v\n", response.GetAverage())
}
