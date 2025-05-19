package main

import (
	"context"
	"io"
	"log"

	pb "github.com/kangs029/grpc-go/calculator/proto"
)

func doPrimes(client pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimeRequest{
		Number: 120,
	}
	streamRes, err := client.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v\n", err)
	}

	for {
		res, err := streamRes.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading from stream: %v\n", err)
		}
		log.Printf("Primes: %v\n", res.PrimeNumber)
	}
}
