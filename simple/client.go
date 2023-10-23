package main

import (
	"context"
	"fmt"
	"simple-grpc/proto"

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

	grpcClient := proto.NewSimpleClient(conn)
	res, err := grpcClient.Route(
		context.Background(),
		&proto.SimpleRequest{
			Data: "Jeremy",
		},
	)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		fmt.Println("CODE:", res.Code, "; VALUE:", res.Value)
	}
}
