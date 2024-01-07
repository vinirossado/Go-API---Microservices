package main

import pb "github.com/rossado/grpc/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
