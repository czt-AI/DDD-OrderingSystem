package grpc

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// GRPCClientRouting gRPC客户端路由
type GRPCClientRouting struct {
	gatewayMux     *runtime.ServeMux
	grpcDialer     grpc.Dialer
	healthEndpoint string
	servers        []string
	roundRobin     int
	mu             sync.Mutex
}

// NewGRPCClientRouting 创建新的gRPC客户端路由实例
func NewGRPCClientRouting(gatewayMux *runtime.ServeMux, grpcDialer grpc.Dialer, healthEndpoint string, servers []string) *GRPCClientRouting {
	return &GRPCClientRouting{
		gatewayMux:     gatewayMux,
		grpcDialer:     grpcDialer,
		healthEndpoint: healthEndpoint,
		servers:        servers,
		roundRobin:     0,
		mu:             sync.Mutex{},
	}
}

// AddServer 添加服务器到负载均衡器
func (s *GRPCClientRouting) AddServer(server string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.servers = append(s.servers, server)
}

// GetServer 获取下一个服务器地址
func (s *GRPCClientRouting) GetServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	server := s.servers[s.roundRobin]
	s.roundRobin = (s.roundRobin + 1) % len(s.servers)
	return server
}

// Start 启动客户端
func (s *GRPCClientRouting) Start() error {
	// 启动gRPC网关服务器
	go func() {
		if err := http.ListenAndServe(s.healthEndpoint, s.gatewayMux); err != nil {
			log.Fatalf("Failed to serve gRPC gateway: %v", err)
		}
	}()

	return nil
}

// Stop 停止客户端
func (s *GRPCClientRouting) Stop() error {
	// 停止gRPC网关服务器
	if err := http.Server{Handler: s.gatewayMux}.Shutdown(); err != nil {
		return err
	}

	return nil
}
