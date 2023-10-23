package main

import (
	"context"
	"fmt"
	"io"
	"server-side-streaming-grpc/proto"
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

	grpcClient := proto.NewStreamServerClient(conn)
	stream, err := grpcClient.ListValue(
		context.Background(),
		&proto.SimpleRequest{
			Data: "Jeremy",
		},
	)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("ERROR:", err.Error())
			}
			fmt.Println(res.StreamValue)
			break
		}

		// can use close send to stop the stream transmission
		// after .CloseSend(), the server won't send further message
		stream.CloseSend()

		time.Sleep(5 * time.Second)
		fmt.Println("Restart transmission after 5 seconds")

		// can use stream.Recv() to restart the transmission
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println("ERROR:", err.Error())
			}
			fmt.Println(res.StreamValue)
		}
	}
}
