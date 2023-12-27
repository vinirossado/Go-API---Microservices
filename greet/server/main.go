package main

import (
	pb "github.com/rossado/grpc/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
