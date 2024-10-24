package grpc

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GRPCClientSecurity gRPC客户端安全
type GRPCClientSecurity struct {
	certificate string
	key         string
}

// NewGRPCClientSecurity 创建新的gRPC客户端安全实例
func NewGRPCClientSecurity(certificate, key string) *GRPCClientSecurity {
	return &GRPCClientSecurity{
		certificate: certificate,
		key:         key,
	}
}

// WrapClientDialer 包装客户端 dialer，用于添加安全
func (s *GRPCClientSecurity) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	// 加载TLS证书和密钥
	creds, err := credentials.NewClientTLSFromFile(s.certificate, s.key)
	if err != nil {
		log.Fatalf("Failed to load TLS credentials %v", err)
	}

	// 使用TLS证书和密钥配置客户端
	return grpc.WithTransportCredentials(creds)(dialer)
}

// RegisterSecurity 注册安全HTTP处理器
func (s *GRPCClientSecurity) RegisterSecurity(handler http.Handler) {
	// 安全配置通常不通过HTTP处理器进行注册，而是直接在gRPC客户端配置中设置
	// 因此，这里不提供HTTP处理器注册方法
}
