package main

import (
	"context"
	pb "github.com/rossado/grpc/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("ListBlog was invoked")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("UError while calling ListBlog: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err != io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}
		log.Println(res)
	}
}
