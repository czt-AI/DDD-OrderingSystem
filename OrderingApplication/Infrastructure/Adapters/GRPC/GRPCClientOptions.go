package grpc

import (
	"google.golang.org/grpc"
)

// GRPCClientOptions gRPC客户端选项
type GRPCClientOptions struct {
	Host       string
	Port       int
	KeepAliveTime time.Duration
	KeepAliveTimeout time.Duration
}

// NewGRPCClientOptions 创建新的gRPC客户端选项实例
func NewGRPCClientOptions(host string, port int, keepAliveTime, keepAliveTimeout time.Duration) *GRPCClientOptions {
	return &GRPCClientOptions{
		Host:       host,
		Port:       port,
		KeepAliveTime: keepAliveTime,
		KeepAliveTimeout: keepAliveTimeout,
	}
}

// ClientOptions 返回gRPC客户端选项
func (o *GRPCClientOptions) ClientOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(
			grpc.WithKeepaliveTime(o.KeepAliveTime),
			grpc.WithKeepaliveTimeout(o.KeepAliveTimeout),
		),
		grpc.WithAddress(fmt.Sprintf("%s:%d", o.Host, o.Port)),
	}
}