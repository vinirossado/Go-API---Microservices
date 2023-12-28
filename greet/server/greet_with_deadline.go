package main

import (
	"context"
	"fmt"
	pb "github.com/rossado/grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("GreetWithDeadline function was invoked with %v\n", request)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + request.FirstName,
	}, nil
}
