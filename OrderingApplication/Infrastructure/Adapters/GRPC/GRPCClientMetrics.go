package grpc

import (
	"context"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// GRPCClientMetrics gRPC客户端指标
type GRPCClientMetrics struct {
	requestDuration *prometheus.Histogram
}

// NewGRPCClientMetrics 创建新的gRPC客户端指标实例
func NewGRPCClientMetrics() *GRPCClientMetrics {
	return &GRPCClientMetrics{
		requestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "grpc_client_request_duration_seconds",
			Help:    "Duration of gRPC client requests in seconds.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}),
	}
}

// WrapClientDialer 包装客户端 dialer，用于添加指标
func (m *GRPCClientMetrics) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)

		m.requestDuration.WithLabelValues(method).Observe(duration.Seconds())

		return err
	})(dialer)
}

// RegisterMetrics 注册指标HTTP处理器
func (m *GRPCClientMetrics) RegisterMetrics(handler http.Handler) {
	prometheus.MustRegister(m.requestDuration)
	handler.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		promhttp.WriteTo(w, promhttp.Handler().ServeHTTP)
	})
}