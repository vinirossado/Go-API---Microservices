package main

import (
	"context"
	pb "github.com/rossado/grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Starting to do a Unary RPC...")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Ross"})
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %s", res.Result)

}
