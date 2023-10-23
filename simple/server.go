package main

import (
	"context"
	"fmt"
	"net"
	"simple-grpc/proto"

	"google.golang.org/grpc"
)

type SimpleService struct {
	proto.UnimplementedSimpleServer
}

func (s *SimpleService) Route(
	ctx context.Context,
	req *proto.SimpleRequest,
) (*proto.SimpleResponse, error) {
	res := proto.SimpleResponse{
		Code:  200,
		Value: "Hello, " + req.Data,
	}
	return &res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen to port 9000 ...")

	server := grpc.NewServer()
	proto.RegisterSimpleServer(server, &SimpleService{})
	server.Serve(listen)
}
