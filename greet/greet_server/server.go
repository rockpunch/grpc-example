package main

import (
	"context"
	"fmt"
	"github.com/rockpunch/grpc-example/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function invoked with %v\n", req)
	fn := req.GetGreeting().GetFirstName()
	ln := req.GetGreeting().GetLastName()
	result := "Hello " + fn + ln
	res := &greetpb.GreetResponse{
		Result:               result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
