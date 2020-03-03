package main

import (
	"context"
	"fmt"
	"github.com/rockpunch/grpc-example/calculator/calc_pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	handleError(err)

	defer conn.Close()

	c := calcpb.NewCalcServiceClient(conn)

	fmt.Println("begin request")

	req := &calcpb.CalcRequest{
		Room:                 &calcpb.Room{
			Length:               440,
			Width:                680,
			Height:               950,
		},
	}

	res, err := c.Calc(context.Background(), req)
	handleError(err)

	log.Printf("LOG\nPrinting Result\n\nAxial: %v\nTangential: %v\nOblique: %v\n", res.Axial, res.Tangential, res.Oblique)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

