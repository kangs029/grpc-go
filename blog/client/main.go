package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kangs029/grpc-go/blog/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)
	id := createBlog(client)
	readBlog(client, id) //valid
	//readBlog(client, "invalidId") //invalid

	updateBlog(client, id)
	listBlog(client)
	//deleteBlog(client, id)
	createBlog(client)
	createBlog(client)

	listBlog(client)

}
