package main

import (
	"context"
	pb "github.com/rossado/grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {

	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Ross",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", e)
		}
	}

	log.Printf("GreetWithDealine: %s\n", res.Result)
}
