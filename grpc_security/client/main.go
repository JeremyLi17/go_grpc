package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/jeremyli17/grpc/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/* When use token-based auth
type ClientTokenAuth struct{}

func (c ClientTokenAuth) GetRequestMetadata(
	ctx context.Context,
	uri ...string,
) (map[string]string, error) {
	return map[string]string{
		"appId":  "testId",
		"appKey": "key",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	// if you want to combine TLS, return true
	return false
}
*/

func main() {
	// creat credential
	creds, _ := credentials.NewClientTLSFromFile(
		"/Users/jeremy/Desktop/Work/Go_Tutorial/grpc/key/test.pem",
		// the second para is the domain name
		// this should be dynamic obtain from the browser in production
		"*.jeremy.com",
	)

	// token-based options
	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	// create connection
	conn, err := grpc.Dial(
		"127.0.0.1:9090",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create client
	client := pb.NewSayHelloClient(conn)

	// send request
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(client pb.SayHelloClient) {
			defer wg.Done()
			rsp, err := client.SayHello(context.Background(), &pb.HelloRequest{
				RequestName: "Jeremy",
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(rsp.ResponseMsg)
			}
		}(client)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("Finish!")
}
