package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// LoggingMiddleware 日志中间件
type LoggingMiddleware struct{}

// Wrap 服务包装函数
func (l *LoggingMiddleware) Wrap(server *grpc.Server) {
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

// HealthCheckMiddleware 健康检查中间件
type HealthCheckMiddleware struct{}

// Wrap 服务包装函数
func (m *HealthCheckMiddleware) Wrap(server *grpc.Server) {
	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(server, &HealthCheckServer{})
}

// HealthCheckServer 健康检查服务器实现
type HealthCheckServer struct{}

// Check 健康检查端点
func (s *HealthCheckServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	// 返回健康状态
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
