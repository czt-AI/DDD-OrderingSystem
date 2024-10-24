package grpc

import (
	"time"

	"google.golang.org/grpc"
)

// GRPCClientConfig gRPC客户端配置
type GRPCClientConfig struct {
	Host       string
	Port       int
	KeepAliveTime time.Duration
	KeepAliveTimeout time.Duration
}

// NewGRPCClientConfig 创建新的gRPC客户端配置实例
func NewGRPCClientConfig(host string, port int, keepAliveTime, keepAliveTimeout time.Duration) *GRPCClientConfig {
	return &GRPCClientConfig{
		Host:       host,
		Port:       port,
		KeepAliveTime: keepAliveTime,
		KeepAliveTimeout: keepAliveTimeout,
	}
}

// ClientOptions 返回gRPC客户端选项
func (c *GRPCClientConfig) ClientOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(
			grpc.WithKeepaliveTime(c.KeepAliveTime),
			grpc.WithKeepaliveTimeout(c.KeepAliveTimeout),
		),
	}
}