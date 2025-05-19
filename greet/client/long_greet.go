package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kangs029/grpc-go/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	// create bunch of messages  send to the server
	requests := []*pb.GreetRequest{
		{FirstName: "Kang"},
		{FirstName: "Soo"},
		{FirstName: "Jin"},
	}

	// create a stream by invoking the LongGreet function
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
		// simulate a delay to see the streaming in action
		time.Sleep(1 * time.Second)
	}

	// close the stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet Response: %v\n", res.Result)
}
