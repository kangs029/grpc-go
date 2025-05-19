package main

import (
	"context"
	"io"
	"log"

	pb "github.com/kangs029/grpc-go/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	//create a request
	req := &pb.GreetRequest{
		FirstName: "Kangs",
	}

	// create a stream by invoking the GreetManyTimes RPC
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v", err)
	}

	// receive a stream of messages from the server
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading from stream: %v", err)
		}

		log.Printf("GreetManyTimes Response: %s", res.Result)
	}
}
