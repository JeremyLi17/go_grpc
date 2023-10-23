package interceptor

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jeremyli17/grpc/server/limiter"
	"google.golang.org/grpc"
)

func LimiterInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// 10 event per second
	if limiter.Limiter.Allow() {
		rsp, err := handler(ctx, req)
		return rsp, err
	} else {
		return nil, errors.New("qps over limit")
	}
}

func LoggerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Println("Received a request at", time.Now())
	// startTime := time.Now()
	rsp, err := handler(ctx, req)

	// fmt.Println("Process time:", time.Since(startTime).Milliseconds(), "ms")
	return rsp, err
}
