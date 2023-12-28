package main

import (
	"context"
	pb "github.com/rossado/grpc/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Creating the blog was invoked")

	blog := &pb.Blog{
		AuthorId: "Ross",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	log.Printf("Blog has been created: %v\n", res)

	return res.Id
}
