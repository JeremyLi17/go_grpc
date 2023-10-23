package main

import (
	"context"
	"net"

	"github.com/jeremyli17/grpc/server/interceptor"
	"github.com/jeremyli17/grpc/server/limiter"
	pb "github.com/jeremyli17/grpc/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloResponse, error) {
	/* When use token-based authentication
	// get metadata
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no token")
	}

	var appId string
	var appKey string

	if v, ok := metadata["appid"]; ok {
		appId = v[0]
	}
	if v, ok := metadata["appkey"]; ok {
		appKey = v[0]
	}

	// verify user
	fmt.Println("appId:", appId)
	fmt.Println("appKey:", appKey)
	*/

	return &pb.HelloResponse{
		ResponseMsg: "hello!" + req.RequestName,
	}, nil
}

func main() {
	// open port
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	creds, _ := credentials.NewServerTLSFromFile(
		"/Users/jeremy/Desktop/Work/Go_Tutorial/grpc/key/test.pem",
		"/Users/jeremy/Desktop/Work/Go_Tutorial/grpc/key/test.key",
	)

	// register a regular grpc server
	// grpcServer := grpc.NewServer()

	// register grpc with TLS & limiter
	limiter.InitLimiter()
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),

		// only 1 interceptor can use:
		// grpc.UnaryInterceptor(interceptor.LimiterInterceptor),

		// more than 1 interceptor:
		grpc.ChainUnaryInterceptor(
			interceptor.LoggerInterceptor,
			interceptor.LimiterInterceptor,
		),
	)
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// start server
	grpcServer.Serve(listen)
}
