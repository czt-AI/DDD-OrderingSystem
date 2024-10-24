package grpc

import (
	"context"
	"DDD-OrderingSystem/Infrastructure/Adapters/HealthCheck"
	"google.golang.org/grpc"
	"healthcheckservice"
)

// HealthCheckServiceServer gRPC健康检查服务服务器
type HealthCheckServiceServer struct{}

// NewHealthCheckServiceServer 创建新的HealthCheckServiceServer实例
func NewHealthCheckServiceServer() *HealthCheckServiceServer {
	return &HealthCheckServiceServer{}
}

// Check 健康检查端点
func (s *HealthCheckServiceServer) Check(ctx context.Context, req *healthcheckservice.CheckRequest) (*healthcheckservice.CheckResponse, error) {
	// 执行健康检查逻辑
	healthStatus, err := HealthCheck.CheckServiceHealth()
	if err != nil {
		return nil, err
	}

	// 返回健康检查结果
	return &healthcheckservice.CheckResponse{
		Status: healthStatus,
	}, nil
}
