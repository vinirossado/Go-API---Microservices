package main

import (
	"context"
	pb "github.com/rossado/grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone() invoked...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Ross"},
		{FirstName: "John"},
		{FirstName: "Paul"},
		{FirstName: "Ringo"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error while receiving response from GreetEveryone: %v", err)
				break
			}

			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
