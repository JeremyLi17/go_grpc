package main

import (
	"both-side-streaming-grpc/proto"
	"context"
	"fmt"
	"io"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		":9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	grpcClient := proto.NewStreamClient(conn)

	stream, err := grpcClient.Conversation(context.Background())
	if err != nil {
		panic(err)
	}

	for n := 1; n <= 5; n++ {
		err := stream.Send(&proto.StreamRequest{
			Question: strconv.Itoa(n) + ". question: xxx",
		})
		if err != nil {
			fmt.Println("ERROR", err.Error())
			break
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR", err.Error())
			break
		}

		fmt.Println("Result back:", res.Answer)
	}
	stream.CloseSend()
}
