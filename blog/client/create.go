package main

import (
	"context"
	"log"

	pb "github.com/kangs029/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Creating a blog...")

	// create a blog
	blog := &pb.Blog{
		AuthorId: "kangs",
		Title:    "My first blog",
		Content:  "This is the content of my first blog",
	}

	// call the CreateBlog RPC
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error creating blog: %v", err)
	}

	log.Printf("Blog created: %s", res.GetId())
	return res.GetId()
}
