package grpc

import (
	"context"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

// GRPCServerTelemetry gRPC服务器遥测
type GRPCServerTelemetry struct {
	requestDuration *prometheus.Histogram
}

// NewGRPCServerTelemetry 创建新的gRPC服务器遥测实例
func NewGRPCServerTelemetry() *GRPCServerTelemetry {
	return &GRPCServerTelemetry{
		requestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds",
			Help:    "Duration of gRPC requests in seconds.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}),
	}
}

// Wrap 服务包装函数
func (t *GRPCServerTelemetry) Wrap(server *grpc.Server) {
	// 添加遥测中间件到每个服务方法
	grpc.ServerMethodMiddleware(server, func(ctx context.Context, fullMethod string, handler grpc.UnaryHandler) grpc.UnaryHandler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			start := time.Now()

			resp, err := handler(ctx, req)
			duration := time.Since(start)

			t.requestDuration.WithLabelValues(fullMethod).Observe(duration.Seconds())

			return resp, err
		}
	})
}

// RegisterMetrics 注册指标HTTP处理器
func (t *GRPCServerTelemetry) RegisterMetrics(handler http.Handler) {
	prometheus.MustRegister(t.requestDuration)
	handler.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		promhttp.WriteTo(w, promhttp.Handler().ServeHTTP)
	})
}
