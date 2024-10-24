package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// GRPCClientLogger gRPC客户端日志记录器
type GRPCClientLogger struct{}

// WrapClientDialer 包装客户端 dialer，用于添加日志
func (l *GRPCClientLogger) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		log.Printf("Sending %s request to %s", method, cc.GetAddr())
		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)
		log.Printf("Received %s response from %s in %s, Error: %v", method, cc.GetAddr(), duration, err)
		return err
	})(dialer)
}
