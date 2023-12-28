package main

import (
	"context"
	pb "github.com/rossado/grpc/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("Reading the blog was invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
