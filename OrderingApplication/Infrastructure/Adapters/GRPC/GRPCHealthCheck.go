package grpc

import (
	"context"
	"DDD-OrderingSystem/OrderingApplication/Infrastructure/Adapters/HealthCheck"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// GRPCHealthCheck gRPC健康检查服务
type GRPCHealthCheck struct{}

// RegisterHealthServer 注册健康检查服务
func (g *GRPCHealthCheck) RegisterHealthServer(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, &HealthCheckServer{})
}

// HealthCheckServer 健康检查服务器实现
type HealthCheckServer struct{}

// Check 健康检查端点
func (s *HealthCheckServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	// 执行健康检查逻辑
	healthStatus, err := HealthCheck.CheckServiceHealth()
	if err != nil {
		return nil, err
	}

	// 返回健康检查结果
	return &grpc_health_v1.HealthCheckResponse{
		Status: healthStatus,
	}, nil
}
