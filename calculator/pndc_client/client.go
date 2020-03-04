package main

import (
	"context"
	"fmt"
	calc "github.com/rockpunch/grpc-example/calculator/calcpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Begin client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := calc.NewPndcServiceClient(conn)

	doServerStreaming(c)
}

func doServerStreaming(c calc.PndcServiceClient) {
	fmt.Println("begin server streaming grpc")

	req := &calc.PndcRequest{
		Number:               332,
	}

	resultStream, err := c.Pndc(context.Background(), req)

	if err != nil {
		log.Fatalf("error getting resultStream: %v\n", err)
	}
	for {
		msg, err:= resultStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error during stream recv: %v\n", err)
		}
		log.Printf("result from calc stream: %v\n", msg.GetResult())
	}
}