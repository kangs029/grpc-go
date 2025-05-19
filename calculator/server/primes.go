package main

import (
	"log"

	pb "github.com/kangs029/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

func (s *server) Primes(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	log.Printf("Primes function was invoked with %v\n", in)

	divisor := int32(2)
	number := in.GetNumber()
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				PrimeNumber: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
