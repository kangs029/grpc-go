package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/kangs029/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

func (s *server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("LongGreet function was invoked")

	res := ""
	for {
		// receive the message from the client
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			return err
		}
		log.Printf("Received request: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.GetFirstName())
	}
}
