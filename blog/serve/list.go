package main

import (
	"context"
	"fmt"
	pb "github.com/rossado/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) ListBlogs(_ *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListBlogs was invoked")

	ctx := context.Background()
	cur, err := collection.Find(ctx, primitive.D{{}})

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v\n", err))
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB: %v\n", err))
		}
		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v\n", err))
	}

	return nil
}
