package main

import (
	"context"
	"fmt"
	pb "github.com/rossado/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (*Server) DeleteBlog(ctx context.Context, req *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked with %v\n", req)

	oid, err := primitive.ObjectIDFromHex(req.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog in MongoDB: %v\n", err),
		)
	}

	return &emptypb.Empty{}, nil
}
