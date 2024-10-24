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

// GRPCServerRouting gRPC服务器路由
type GRPCServerRouting struct {
	grpcServer     *grpc.Server
	gatewayMux     *runtime.ServeMux
	grpcListener   net.Listener
	healthListener net.Listener
	roundRobin     int
	servers        []string
	mu             sync.Mutex
}

// NewGRPCServerRouting 创建新的gRPC服务器路由实例
func NewGRPCServerRouting(grpcServer *grpc.Server, gatewayMux *runtime.ServeMux, grpcListener net.Listener, healthListener net.Listener) *GRPCServerRouting {
	return &GRPCServerRouting{
		grpcServer:     grpcServer,
		gatewayMux:     gatewayMux,
		grpcListener:   grpcListener,
		healthListener: healthListener,
		roundRobin:     0,
		servers:        []string{},
		mu:             sync.Mutex{},
	}
}

// AddServer 添加服务器到负载均衡器
func (s *GRPCServerRouting) AddServer(server string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.servers = append(s.servers, server)
}

// GetServer 获取下一个服务器地址
func (s *GRPCServerRouting) GetServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	server := s.servers[s.roundRobin]
	s.roundRobin = (s.roundRobin + 1) % len(s.servers)
	return server
}

// Start 启动服务器
func (s *GRPCServerRouting) Start() error {
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
func (s *GRPCServerRouting) Stop() error {
	// 停止gRPC服务器
	s.grpcServer.Stop()

	// 停止gRPC网关服务器
	if err := http.Server{Handler: s.gatewayMux}.Shutdown(); err != nil {
		return err
	}

	// 停止健康检查服务器
	return http.Serve(s.healthListener, nil)
}
