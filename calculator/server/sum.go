package main

import (
	"context"
	pb "github.com/rossado/grpc/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", req)
	return &pb.SumResponse{Result: req.FirstNumber + req.SecondNumber}, nil
}
