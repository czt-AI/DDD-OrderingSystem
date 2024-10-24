package grpc

import (
	"time"

	"google.golang.org/grpc"
)

// GRPCServerConfig gRPC服务器配置
type GRPCServerConfig struct {
	Port       int
	KeepAliveTime time.Duration
	KeepAliveTimeout time.Duration
}

// NewGRPCServerConfig 创建新的gRPC服务器配置实例
func NewGRPCServerConfig(port int, keepAliveTime, keepAliveTimeout time.Duration) *GRPCServerConfig {
	return &GRPCServerConfig{
		Port:       port,
		KeepAliveTime: keepAliveTime,
		KeepAliveTimeout: keepAliveTimeout,
	}
}

// ServerOptions 返回gRPC服务器选项
func (c *GRPCServerConfig) ServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.KeepaliveMinTime(c.KeepAliveTime),
		grpc.KeepaliveTimeout(c.KeepAliveTimeout),
	}
}
