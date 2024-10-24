package grpc

import (
	"context"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GRPCServerSecurity gRPC服务器安全
type GRPCServerSecurity struct {
	certificate string
	key         string
}

// NewGRPCServerSecurity 创建新的gRPC服务器安全实例
func NewGRPCServerSecurity(certificate, key string) *GRPCServerSecurity {
	return &GRPCServerSecurity{
		certificate: certificate,
		key:         key,
	}
}

// Wrap 服务包装函数
func (s *GRPCServerSecurity) Wrap(server *grpc.Server) {
	// 加载TLS证书和密钥
	creds, err := credentials.NewServerTLSFromFile(s.certificate, s.key)
	if err != nil {
		log.Fatalf("Failed to load TLS credentials %v", err)
	}

	// 使用TLS证书和密钥配置服务器
	server.Creds = creds
}

// RegisterSecurity 注册安全HTTP处理器
func (s *GRPCServerSecurity) RegisterSecurity(handler http.Handler) {
	// 安全配置通常不通过HTTP处理器进行注册，而是直接在gRPC服务器配置中设置
	// 因此，这里不提供HTTP处理器注册方法
}
