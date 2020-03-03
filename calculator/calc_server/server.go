package main

import (
	"context"
	"fmt"
	calcpb "github.com/rockpunch/grpc-example/calculator/calc_pb"
	rmcalc "github.com/rockpunch/grpc-example/calculator/rm_calc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}


func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	handleError(err)

	s := grpc.NewServer()

	calcpb.RegisterCalcServiceServer(s, &server{})

	err = s.Serve(lis); if err != nil {
		log.Fatalf("cannot serve: %v\n", err)
	}
}

func (*server) Calc (ctx context.Context, request *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("Calc function invoked with: %v\n", request)
	l := request.Room.GetLength()
	w := request.Room.GetWidth()
	h := request.Room.GetHeight()
	axial, err := rmcalc.CalcAxial(l, w, h)
	tangential, err := rmcalc.CalcTangential(l, w, h)
	oblique, err := rmcalc.CalcOblique(l, w, h)
	handleError(err)

	response := &calcpb.CalcResponse{
		Axial:                axial,
		Tangential:           tangential,
		Oblique:              oblique,
	}
	return response, err
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}