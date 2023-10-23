package main

import (
	"client-side-streaming-grpc/proto"
	"fmt"
	"io"
	"net"

	"google.golang.org/grpc"
)

type SimpleService struct {
	proto.UnimplementedSimpleServerServer
}

func (s *SimpleService) RouteList(
	srv proto.SimpleServer_RouteListServer,
) error {
	for {
		// get data from stream
		res, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&proto.SimpleResponse{
				Code:  1,
				Value: "Okay",
			})
		}
		if err != nil {
			fmt.Println("ERROR:", err.Error())
		}
		fmt.Println(res.StreamValue)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	fmt.Println("Listen to port 9000 ...")

	server := grpc.NewServer()
	proto.RegisterSimpleServerServer(server, &SimpleService{})
	server.Serve(listen)
}
