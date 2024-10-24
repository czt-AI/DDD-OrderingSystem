package grpc

import (
	"google.golang.org/grpc"
)

// GRPCServerConfiguration gRPC服务器配置
type GRPCServerConfiguration struct {
	Host       string
	Port       int
	KeepAliveTime time.Duration
	KeepAliveTimeout time.Duration
}

// NewGRPCServerConfiguration 创建新的gRPC服务器配置实例
func NewGRPCServerConfiguration(host string, port int, keepAliveTime, keepAliveTimeout time.Duration) *GRPCServerConfiguration {
	return &GRPCServerConfiguration{
		Host:       host,
		Port:       port,
		KeepAliveTime: keepAliveTime,
		KeepAliveTimeout: keepAliveTimeout,
	}
}

// ServerOptions 返回gRPC服务器选项
func (c *GRPCServerConfiguration) ServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.BindAddr(fmt.Sprintf("%s:%d", c.Host, c.Port)),
		grpc.KeepaliveMinTime(c.KeepAliveTime),
		grpc.KeepaliveTimeout(c.KeepAliveTimeout),
	}
}
