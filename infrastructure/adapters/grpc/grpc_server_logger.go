package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// GRPCServerLogger gRPC服务器日志记录器
type GRPCServerLogger struct{}

// Wrap 服务包装函数
func (l *GRPCServerLogger) Wrap(server *grpc.Server) {
	// 添加日志中间件到每个服务方法
	grpc.ServerMethodMiddleware(server, func(ctx context.Context, fullMethod string, handler grpc.UnaryHandler) grpc.UnaryHandler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			start := time.Now()
			log.Printf("Received %s request", fullMethod)
			resp, err := handler(ctx, req)
			duration := time.Since(start)
			log.Printf("Handled %s request in %s, Error: %v", fullMethod, duration, err)
			return resp, err
		}
	})
}
