package main

import (
	"context"
	pb "github.com/rossado/grpc/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked..")
	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 1, SecondNumber: 2})
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Sum: %d", res.Result)

}
