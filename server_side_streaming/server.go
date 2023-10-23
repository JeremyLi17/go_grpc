package main

import (
	"fmt"
	"net"
	"server-side-streaming-grpc/proto"
	"strconv"

	"google.golang.org/grpc"
)

type StreamService struct {
	proto.UnimplementedStreamServerServer
}

func (s *StreamService) ListValue(
	req *proto.SimpleRequest,
	srv proto.StreamServer_ListValueServer,
) error {
	for i := 0; i < 5; i++ {
		// maximum length of each response is math.MaxInt32 bytes
		err := srv.Send(&proto.StreamResponse{
			StreamValue: "hello" + req.Data + " " + strconv.Itoa(i),
		})
		if err != nil {
			return err
		}

		fmt.Println("Sent:", strconv.Itoa(i))
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen to port 9000 ...")

	server := grpc.NewServer()
	proto.RegisterStreamServerServer(server, &StreamService{})
	server.Serve(listen)
}
