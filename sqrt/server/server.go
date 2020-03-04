package main

import (
	"context"
	"fmt"
	"github.com/rockpunch/grpc-example/sqrt/sqrtpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type server struct{}

func (s *server) Calc(ctx context.Context, req *sqrtpb.SqrtRequest) (*sqrtpb.SqrtResponse, error) {
	num := req.GetNumber()
	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v\n", num),
			)
	}
	return &sqrtpb.SqrtResponse{
		Sqrt:                 num,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	sqrtpb.RegisterCalcSqrtServer(s, &server{})

	err = s.Serve(lis); if err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}


