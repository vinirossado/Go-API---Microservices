package main

import (
	pb "github.com/rossado/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)
}
