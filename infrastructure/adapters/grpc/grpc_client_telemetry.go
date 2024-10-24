package grpc

import (
	"context"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// GRPCClientTelemetry gRPC客户端遥测
type GRPCClientTelemetry struct {
	requestDuration *prometheus.Histogram
}

// NewGRPCClientTelemetry 创建新的gRPC客户端遥测实例
func NewGRPCClientTelemetry() *GRPCClientTelemetry {
	return &GRPCClientTelemetry{
		requestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "grpc_client_request_duration_seconds",
			Help:    "Duration of gRPC client requests in seconds.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}),
	}
}

// WrapClientDialer 包装客户端 dialer，用于添加遥测
func (t *GRPCClientTelemetry) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)

		t.requestDuration.WithLabelValues(method).Observe(duration.Seconds())

		return err
	})(dialer)
}

// RegisterMetrics 注册指标HTTP处理器
func (t *GRPCClientTelemetry) RegisterMetrics(handler http.Handler) {
	prometheus.MustRegister(t.requestDuration)
	handler.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		promhttp.WriteTo(w, promhttp.Handler().ServeHTTP)
	})
}
