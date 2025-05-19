package main

import (
	"io"
	"log"

	pb "github.com/kangs029/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

func (s *server) Average(stream grpc.ClientStreamingServer[pb.AverageRequest, pb.AverageResponse]) error {
	log.Printf("Average function was invoked ")
	sum := int32(0)
	count := 0

	for {
		in, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Average: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream, %v", err)
		}

		log.Printf("Receiving number: %v\n", in.GetNumber())
		sum += in.GetNumber()
		count++
	}
}
