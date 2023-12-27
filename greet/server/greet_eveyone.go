package main

import (
	pb "github.com/rossado/grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked with a streaming request")
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we have finished reading the client stream
				return nil
			}
			log.Fatalf("Error while reading client stream: %v", err)
		}

		res := "Hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}

	}
	return nil
}
