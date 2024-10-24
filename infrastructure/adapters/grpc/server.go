package grpc

import (
	"context"
	"net"
	"net/http"

	"DDD-OrderingSystem/Infrastructure/Adapters/HealthCheck"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Server gRPC服务器
type Server struct {
	grpcServer     *grpc.Server
	gatewayMux     *runtime.ServeMux
	grpcListener   net.Listener
	healthServer   *grpc_health_v1.Server
	healthListener net.Listener
}

// NewServer 创建新的gRPC服务器实例
func NewServer() (*Server, error) {
	// 创建gRPC服务器
	grpcServer := grpc.NewServer()
	healthServer := grpc_health_v1.NewServer()

	// 注册服务
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	// 创建gRPC网关服务
	gatewayMux := runtime.NewServeMux()
	endpoint := ":8080"
	healthEndpoint := ":8081"

	// 设置gRPC网关服务
	err := RegisterGRPCGatewayServer(gatewayMux, endpoint, grpcServer)
	if err != nil {
		return nil, err
	}

	// 创建gRPC监听器
	grpcListener, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	// 创建健康检查监听器
	healthListener, err := net.Listen("tcp", healthEndpoint)
	if err != nil {
		return nil, err
	}

	return &Server{
		grpcServer:     grpcServer,
		gatewayMux:     gatewayMux,
		grpcListener:   grpcListener,
		healthServer:   healthServer,
		healthListener: healthListener,
	}, nil
}

// Start 启动服务器
func (s *Server) Start() error {
	// 启动gRPC服务器
	go func() {
		if err := s.grpcServer.Serve(s.grpcListener); err != nil {
			logrus.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// 启动gRPC网关服务器
	go func() {
		if err := http.ListenAndServe(s.grpcListener.Addr().String(), s.gatewayMux); err != nil {
			logrus.Fatalf("Failed to serve gRPC gateway: %v", err)
		}
	}()

	// 启动健康检查服务器
	go func() {
		if err := s.healthServer.Serve(s.healthListener); err != nil {
			logrus.Fatalf("Failed to serve health server: %v", err)
		}
	}()

	return nil
}

// Stop 停止服务器
func (s *Server) Stop() error {
	// 停止gRPC服务器
	s.grpcServer.Stop()

	// 停止gRPC网关服务器
	if err := http.Server{Handler: s.gatewayMux}.Shutdown(); err != nil {
		return err
	}

	// 停止健康检查服务器
	s.healthServer.Stop()

	return nil
}

// RegisterGRPCGatewayServer 注册gRPC网关服务器
func RegisterGRPCGatewayServer(mux *runtime.ServeMux, endpoint string, server *grpc.Server) error {
	// 设置gRPC网关选项
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}

	// 注册gRPC网关服务
	return runtime.ServeFromEndpoint(mux, endpoint, server, opts)
}
