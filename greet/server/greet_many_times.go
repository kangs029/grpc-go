package main

import (
	"fmt"
	"log"

	pb "github.com/kangs029/grpc-go/greet/proto"
	"google.golang.org/grpc"
)

func (s *server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.GetFirstName(), i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
