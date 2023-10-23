package main

import (
	"both-side-streaming-grpc/proto"
	"fmt"
	"io"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type StreamService struct {
	proto.UnimplementedStreamServer
}

func (s *StreamService) Conversation(
	srv proto.Stream_ConversationServer,
) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Println("Received:", req.Question)
		err = srv.Send(&proto.StreamResponse{
			Answer: "Answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
	}
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen to port 9000 ...")

	grpcServer := grpc.NewServer()
	proto.RegisterStreamServer(grpcServer, &StreamService{})
	grpcServer.Serve(listen)
}
