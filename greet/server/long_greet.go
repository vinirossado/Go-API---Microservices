package main

import (
	"fmt"
	pb "github.com/rossado/grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")
	res := ""
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we have finished reading the client stream
				return stream.SendAndClose(&pb.GreetResponse{Result: res})
			}
			log.Fatalf("Error while reading client stream: %v", err)
		}

		res += fmt.Sprintf("Hello %s!\n ", req.FirstName)
	}
	return nil
}
