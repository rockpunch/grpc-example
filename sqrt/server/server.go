package main

import (
	"github.com/rockpunch/grpc-example/sqrt/sqrtpb"
	"log"
	"net"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}


}


