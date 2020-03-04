package main

import (
	"context"
	"fmt"
	"github.com/rockpunch/grpc-example/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstname + " " + lastname + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result:               result,
		}
		err := stream.Send(res)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	fmt.Println("Srv run")

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
