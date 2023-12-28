package main

import (
	pb "github.com/rossado/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

var addr = "localhost:50051"

func main() {
	tls := true // change that to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	//doGreetWithDeadline(c, 5*time.Second)
}
