package main

import (
	"context"
	pb "github.com/rossado/grpc/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet() invoked...")

	reqs := []*pb.GreetRequest{
		{FirstName: "Ross"},
		{FirstName: "John"},
		{FirstName: "Paul"},
		{FirstName: "Ringo"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	log.Printf("LongGreet Response: %v\n", res.Result)

}
