package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// GRPCServerInterceptor gRPC服务器拦截器
type GRPCServerInterceptor struct{}

// Wrap 服务包装函数
func (i *GRPCServerInterceptor) Wrap(server *grpc.Server) {
	// 添加拦截器到gRPC服务器
	grpcServerInterceptors := server.Interceptors()
	grpcServerInterceptors = append(grpcServerInterceptors, i)
	server.SetInterceptors(grpcServerInterceptors)
}

// InterceptorFunc 拦截器函数
func (i *GRPCServerInterceptor) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Printf("Handling request: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	duration := time.Since(start)
	log.Printf("Request handled in %s, Error: %v", duration, err)
	return resp, err
}
