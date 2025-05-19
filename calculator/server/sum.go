package main

import (
	"context"
	"log"

	pb "github.com/kangs029/grpc-go/calculator/proto"
)

func (s *server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was involked with : %v\n", req)
	sum := req.GetFirstNumber() + req.GetSecondNumber()
	res := &pb.SumResponse{
		Sum: sum,
	}
	return res, nil
}
