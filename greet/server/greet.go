package main

import (
	"context"
	pb "github.com/rossado/grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", req)
	return &pb.GreetResponse{Result: "Hello " + req.FirstName}, nil
}
