package grpc

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// GRPCClientStats gRPC客户端统计
type GRPCClientStats struct {
	requestCount *prometheus.Counter
}

// NewGRPCClientStats 创建新的gRPC客户端统计实例
func NewGRPCClientStats() *GRPCClientStats {
	return &GRPCClientStats{
		requestCount: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_client_request_count",
			Help: "Total number of gRPC client requests.",
		}),
	}
}

// WrapClientDialer 包装客户端 dialer，用于添加统计
func (s *GRPCClientStats) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		s.requestCount.Inc()
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)

		log.Printf("Sent %s request to %s in %s, Error: %v", method, cc.GetAddr(), duration, err)
		return err
	})(dialer)
}

// RegisterStats 注册统计HTTP处理器
func (s *GRPCClientStats) RegisterStats(handler http.Handler) {
	prometheus.MustRegister(s.requestCount)
	handler.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		promhttp.WriteTo(w, promhttp.Handler().ServeHTTP)
	})
}