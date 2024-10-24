package grpc

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// GRPCClientRouting gRPC客户端路由
type GRPCClientRouting struct {
	gatewayMux     *runtime.ServeMux
	grpcDialer     grpc.Dialer
	healthEndpoint string
}

// NewGRPCClientRouting 创建新的gRPC客户端路由实例
func NewGRPCClientRouting(gatewayMux *runtime.ServeMux, grpcDialer grpc.Dialer, healthEndpoint string) *GRPCClientRouting {
	return &GRPCClientRouting{
		gatewayMux:     gatewayMux,
		grpcDialer:     grpcDialer,
		healthEndpoint: healthEndpoint,
	}
}

// Start 启动路由
func (r *GRPCClientRouting) Start() error {
	// 启动gRPC网关服务器
	go func() {
		if err := http.ListenAndServe(r.healthEndpoint, r.gatewayMux); err != nil {
			log.Fatalf("Failed to serve gRPC gateway: %v", err)
		}
	}()

	return nil
}

// Stop 停止路由
func (r *GRPCClientRouting) Stop() error {
	// 停止gRPC网关服务器
	if err := http.Server{Handler: r.gatewayMux}.Shutdown(); err != nil {
		return err
	}

	return nil
}
