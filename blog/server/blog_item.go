package main

import (
	pb "github.com/kangs029/grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

// helper function to convert a BlogItem to a Blog
func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(), //Hex is doing is that basically is taking this ID and turn it into a string
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}
}
