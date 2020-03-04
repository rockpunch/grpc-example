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
	"time"
)

type server struct{}

func (s *server) CalcWithDeadline(ctx context.Context, req *sqrtpb.SqrtRequestWithDeadline) (*sqrtpb.SqrtResponseWithDeadline, error) {
	fmt.Printf("CalcWithDeadline invoked with %v\n", req)
	num := req.GetNumber()

	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v\n", num),
		)
	}

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("the client cancelled the request")
			return nil, status.Error(codes.DeadlineExceeded, "the client cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}

	return &sqrtpb.SqrtResponseWithDeadline{
		Root: num,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	sqrtpb.RegisterCalcSqrtServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
