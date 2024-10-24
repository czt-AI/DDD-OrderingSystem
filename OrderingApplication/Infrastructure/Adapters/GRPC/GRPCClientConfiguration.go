package grpc

import (
	"google.golang.org/grpc"
)

// GRPCClientConfiguration gRPC客户端配置
type GRPCClientConfiguration struct {
	Host       string
	Port       int
	KeepAliveTime time.Duration
	KeepAliveTimeout time.Duration
}

// NewGRPCClientConfiguration 创建新的gRPC客户端配置实例
func NewGRPCClientConfiguration(host string, port int, keepAliveTime, keepAliveTimeout time.Duration) *GRPCClientConfiguration {
	return &GRPCClientConfiguration{
		Host:       host,
		Port:       port,
		KeepAliveTime: keepAliveTime,
		KeepAliveTimeout: keepAliveTimeout,
	}
}

// ClientOptions 返回gRPC客户端选项
func (c *GRPCClientConfiguration) ClientOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(
			grpc.WithKeepaliveTime(c.KeepAliveTime),
			grpc.WithKeepaliveTimeout(c.KeepAliveTimeout),
		),
		grpc.WithAddress(fmt.Sprintf("%s:%d", c.Host, c.Port)),
	}
}