package grpc

import (
	"google.golang.org/grpc"
)

// GRPCServerOptions gRPC服务器选项
type GRPCServerOptions struct {
	Port       int
	KeepAliveTime time.Duration
	KeepAliveTimeout time.Duration
}

// NewGRPCServerOptions 创建新的gRPC服务器选项实例
func NewGRPCServerOptions(port int, keepAliveTime, keepAliveTimeout time.Duration) *GRPCServerOptions {
	return &GRPCServerOptions{
		Port:       port,
		KeepAliveTime: keepAliveTime,
		KeepAliveTimeout: keepAliveTimeout,
	}
}

// ServerOptions 返回gRPC服务器选项
func (o *GRPCServerOptions) ServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.BindAddr(fmt.Sprintf(":%d", o.Port)),
		grpc.KeepaliveMinTime(o.KeepAliveTime),
		grpc.KeepaliveTimeout(o.KeepAliveTimeout),
	}
}
