package grpc

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// GRPCServerLoadBalancing gRPC服务器负载均衡
type GRPCServerLoadBalancing struct {
	grpcServer     *grpc.Server
	gatewayMux     *runtime.ServeMux
	grpcListener   net.Listener
	healthListener net.Listener
	roundRobin     int
	servers        []string
	mu             sync.Mutex
}

// NewGRPCServerLoadBalancing 创建新的gRPC服务器负载均衡实例
func NewGRPCServerLoadBalancing(servers []string) *GRPCServerLoadBalancing {
	return &GRPCServerLoadBalancing{
		servers: servers,
		roundRobin: 0,
	}
}

// AddServer 添加服务器到负载均衡器
func (s *GRPCServerLoadBalancing) AddServer(server string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.servers = append(s.servers, server)
}

// GetServer 获取下一个服务器地址
func (s *GRPCServerLoadBalancing) GetServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	server := s.servers[s.roundRobin]
	s.roundRobin = (s.roundRobin + 1) % len(s.servers)
	return server
}

// Start 启动服务器
func (s *GRPCServerLoadBalancing) Start() error {
	// 启动gRPC服务器
	go func() {
		if err := s.grpcServer.Serve(s.grpcListener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// 启动gRPC网关服务器
	go func() {
		if err := http.ListenAndServe(s.grpcListener.Addr().String(), s.gatewayMux); err != nil {
			log.Fatalf("Failed to serve gRPC gateway: %v", err)
		}
	}()

	// 启动健康检查服务器
	go func() {
		if err := http.Serve(s.healthListener, nil); err != nil {
			log.Fatalf("Failed to serve health server: %v", err)
		}
	}()

	return nil
}

// Stop 停止服务器
func (s *GRPCServerLoadBalancing) Stop() error {
	// 停止gRPC服务器
	s.grpcServer.Stop()

	// 停止gRPC网关服务器
	if err := http.Server{Handler: s.gatewayMux}.Shutdown(); err != nil {
		return err
	}

	// 停止健康检查服务器
	return http.Serve(s.healthListener, nil)
}
