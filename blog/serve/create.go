package main

import (
	"context"
	"fmt"
	pb "github.com/rossado/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateBlog(ctx context.Context, req *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", req)

	data := BlogItem{
		AuthorId: req.AuthorId,
		Title:    req.Title,
		Content:  req.Content,
	}
	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %v\n", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID\n"))
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
