package main

import (
	"client-side-streaming-grpc/proto"
	"context"
	"fmt"
	"strconv"
	"time"

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
	defer conn.Close()

	grpcClient := proto.NewSimpleServerClient(conn)
	stream, err := grpcClient.RouteList(context.Background())

	for i := 0; i < 5; i++ {
		stream.Send(&proto.StreamRequest{
			StreamValue: "Stream client " + strconv.Itoa(i),
		})
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		fmt.Println("CODE:", res.Code)
		fmt.Println("VALUE:", res.Value)
	}
}
