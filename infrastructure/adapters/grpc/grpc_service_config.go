package grpc

import (
	"google.golang.org/grpc"
)

// GRPCServiceConfig gRPC服务配置
type GRPCServiceConfig struct {
	Host       string
	Port       int
	MaxConnections int
	KeepAlive   time.Duration
}

// NewGRPCServiceConfig 创建新的gRPC服务配置实例
func NewGRPCServiceConfig(host string, port int, maxConnections int, keepAlive time.Duration) *GRPCServiceConfig {
	return &GRPCServiceConfig{
		Host:       host,
		Port:       port,
		MaxConnections: maxConnections,
		KeepAlive:   keepAlive,
	}
}

// ServerOptions 返回gRPC服务器选项
func (c *GRPCServiceConfig) ServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.MaxConnections(c.MaxConnections),
		grpc.KeepaliveMinTime(c.KeepAlive),
		grpc.KeepaliveTimeout(c.KeepAlive),
	}
}
