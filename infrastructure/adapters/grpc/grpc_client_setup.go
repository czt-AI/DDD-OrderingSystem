package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GRPCClientSetup gRPC客户端设置
type GRPCClientSetup struct {
	host       string
	port       int
	keepAliveTime time.Duration
	keepAliveTimeout time.Duration
}

// NewGRPCClientSetup 创建新的gRPC客户端设置实例
func NewGRPCClientSetup(host string, port int, keepAliveTime, keepAliveTimeout time.Duration) *GRPCClientSetup {
	return &GRPCClientSetup{
		host:       host,
		port:       port,
		keepAliveTime: keepAliveTime,
		keepAliveTimeout: keepAliveTimeout,
	}
}

// Setup 配置gRPC客户端
func (s *GRPCClientSetup) Setup() (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%d", s.host, s.port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(
			grpc.WithKeepaliveTime(s.keepAliveTime),
			grpc.WithKeepaliveTimeout(s.keepAliveTimeout),
		),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
