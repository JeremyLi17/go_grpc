package main

import (
	"context"
	"fmt"
	proto "grpc-gateway/proto/echo"
	"net"
	"time"

	"google.golang.org/grpc"
)

type EchoService struct {
	proto.UnimplementedEchoServer
}

func (s *EchoService) Echo(
	ctx context.Context,
	req *proto.StringMessage,
) (*proto.StringMessage, error) {
	fmt.Println("Receive Request")
	rsp := &proto.StringMessage{
		Value: "Received:" + req.Value + " at " + time.Now().String(),
	}
	return rsp, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterEchoServer(server, &EchoService{})
	server.Serve(listen)
}
