package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/rossado/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) ReadBlog(ctx context.Context, req *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v\n", req)

	oid, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return documentToBlog(data), nil
}
