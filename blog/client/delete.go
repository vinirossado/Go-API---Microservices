package main

import (
	"context"
	pb "github.com/rossado/grpc/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("Deleting the blog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Println("Blog was deleted")
}
