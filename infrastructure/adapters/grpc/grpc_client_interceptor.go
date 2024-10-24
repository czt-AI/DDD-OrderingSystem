package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// GRPCClientInterceptor gRPC客户端拦截器
type GRPCClientInterceptor struct{}

// WrapClientDialer 包装客户端 dialer，用于添加拦截器
func (i *GRPCClientInterceptor) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		log.Printf("Sending request to %s", method)
		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)
		log.Printf("Received response from %s in %s, Error: %v", method, duration, err)
		return err
	})(dialer)
}
