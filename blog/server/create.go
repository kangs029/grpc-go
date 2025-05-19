package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kangs029/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog function was invoked with %v\n", in)

	/*
		notice here that we are not passing the ID and this is because we don't need it here.
		What will happen is that after inserting it in MongoDB, MongoDB will return automatically a new unique ID.
	*/
	blog := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	// Insert the blog into the database
	res, err := collection.InsertOne(context.Background(), blog)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to insert blog: %v", err))
	}

	oId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Failed to convert inserted ID to ObjectID")
	}

	return &pb.BlogId{Id: oId.Hex()}, nil
}
