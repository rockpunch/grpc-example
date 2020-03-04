package main

import (
	"context"
	"fmt"
	"github.com/rockpunch/grpc-example/sqrt/sqrtpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn not connect: %v", err)
	}
	defer conn.Close()

	c := sqrtpb.NewCalcSqrtClient(conn)

	doSqrtWithDeadline(c,5, time.Second * 5)
	doSqrtWithDeadline(c,3, time.Second * 1)

}

func doSqrtWithDeadline(c sqrtpb.CalcSqrtClient, number int32, seconds time.Duration) {
	fmt.Println("begin SqrtWithDeadline rpc")

	// Wait for 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), seconds)
	response, err := c.CalcWithDeadline(ctx, &sqrtpb.SqrtRequestWithDeadline{Number: number})
	defer cancel()

	// Error Handling
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// check if err is code of DeadlineExceeded
			if respErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout: deadline exceeded")
				return
			} else if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We sent a negative number")
				return
			} else {
				fmt.Printf("unexpected error: %v", err)
				return
			}
		} else {
			log.Fatalf("error calling sqrt: %v", err)
			return
		}
	}
	fmt.Printf("result of sqrt of number %v: %v\n", number, response.GetRoot())
}
