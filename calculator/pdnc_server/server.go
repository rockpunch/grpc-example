package main

import (
	"fmt"
	calc "github.com/rockpunch/grpc-example/calculator/calcpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	calc.RegisterPndcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}

func (*server) Pndc(req *calc.PndcRequest, stream calc.PndcService_PndcServer) error {
	fmt.Printf("Pndc function invoked with %v\n", req)

	v := req.GetNumber()
	k := int32(2)

	for {
		if v <= 1 {
			break
		}
		if v % k == 0 {
			res := &calc.PndcResponse{
				Result:               k,
			}
			err := stream.Send(res)
			if err != nil {
				log.Fatalln(err)
			}
			v = v / k
		} else {
			k++
		}
	}
	return nil
}