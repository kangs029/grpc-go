package main

import (
	"context"
	"log"

	pb "github.com/kangs029/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Printf("Blog was deleted: %s", id)
}
