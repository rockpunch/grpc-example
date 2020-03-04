package main

import (
	"context"
	"fmt"
	"github.com/rockpunch/grpc-example/sqrt/sqrtpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn not connect: %v", err)
	}
	defer conn.Close()

	c := sqrtpb.NewCalcSqrtClient(conn)

	doSqrt(c, 10)
	doSqrt(c, -2)
}

func doSqrt(c sqrtpb.CalcSqrtClient, number int32) {
	fmt.Println("begin do sqrt unary rpc")
	response, err := c.Calc(context.Background(), &sqrtpb.SqrtRequest{Number: number})

	// Error Handling
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual err from gRPC (as comma ok idiom)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We sent a negative number")
				return
			}
		} else {
			log.Fatalf("error calling sqrt: %v", err)
			return
		}
	}

	fmt.Printf("result of sqrt of number %v: %v\n", number, response.GetSqrt())
}
