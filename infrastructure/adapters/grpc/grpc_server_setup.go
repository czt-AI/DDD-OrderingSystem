package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GRPCServerSetup gRPC服务器设置
type GRPCServerSetup struct {
	port int
}

// NewGRPCServerSetup 创建新的gRPC服务器设置实例
func NewGRPCServerSetup(port int) *GRPCServerSetup {
	return &GRPCServerSetup{
		port: port,
	}
}

// Start 启动gRPC服务器
func (s *GRPCServerSetup) Start(server *grpc.Server) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	log.Printf("Starting gRPC server on port %d", s.port)
	return server.Serve(listener)
}

// Stop 停止gRPC服务器
func (s *GRPCServerSetup) Stop(server *grpc.Server) error {
	log.Printf("Stopping gRPC server")
	server.Stop()
	return nil
}
