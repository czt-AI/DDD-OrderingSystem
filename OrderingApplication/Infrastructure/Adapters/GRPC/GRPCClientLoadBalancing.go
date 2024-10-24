package grpc

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// GRPCClientLoadBalancing gRPC客户端负载均衡
type GRPCClientLoadBalancing struct {
	grpcDialer     grpc.Dialer
	healthEndpoint string
	servers        []string
	roundRobin     int
	mu             sync.Mutex
}

// NewGRPCClientLoadBalancing 创建新的gRPC客户端负载均衡实例
func NewGRPCClientLoadBalancing(servers []string) *GRPCClientLoadBalancing {
	return &GRPCClientLoadBalancing{
		servers: servers,
		roundRobin: 0,
	}
}

// AddServer 添加服务器到负载均衡器
func (s *GRPCClientLoadBalancing) AddServer(server string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.servers = append(s.servers, server)
}

// GetServer 获取下一个服务器地址
func (s *GRPCClientLoadBalancing) GetServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	server := s.servers[s.roundRobin]
	s.roundRobin = (s.roundRobin + 1) % len(s.servers)
	return server
}

// Start 启动客户端
func (s *GRPCClientLoadBalancing) Start() error {
	// 启动gRPC网关服务器
	go func() {
		if err := http.ListenAndServe(s.healthEndpoint, nil); err != nil {
			log.Fatalf("Failed to serve gRPC gateway: %v", err)
		}
	}()

	return nil
}

// Stop 停止客户端
func (s *GRPCClientLoadBalancing) Stop() error {
	// 停止gRPC网关服务器
	if err := http.Server{Handler: nil}.Shutdown(); err != nil {
		return err
	}

	return nil
}
