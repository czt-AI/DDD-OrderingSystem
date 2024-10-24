package grpc

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

// GRPCServerStats gRPC服务器统计
type GRPCServerStats struct {
	requestCount     *prometheus.Counter
	requestDuration  *prometheus.Histogram
	keepAliveCounter *prometheus.Counter
}

// NewGRPCServerStats 创建新的gRPC服务器统计实例
func NewGRPCServerStats() *GRPCServerStats {
	return &GRPCServerStats{
		requestCount: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_server_request_count",
			Help: "Total number of gRPC server requests.",
		}),
		requestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "grpc_server_request_duration_seconds",
			Help:    "Duration of gRPC server requests in seconds.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}),
		keepAliveCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_server_keepalive_count",
			Help: "Total number of keepalive pings sent by the server.",
		}),
	}
}

// Wrap 服务包装函数
func (s *GRPCServerStats) Wrap(server *grpc.Server) {
	// 添加统计中间件到每个服务方法
	grpc.ServerMethodMiddleware(server, func(ctx context.Context, fullMethod string, handler grpc.UnaryHandler) grpc.UnaryHandler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			s.requestCount.Inc()
			start := time.Now()

			resp, err := handler(ctx, req)
			duration := time.Since(start)

			log.Printf("Handled %s request in %s, Error: %v", fullMethod, duration, err)
			return resp, err
		}
	})
}

// RegisterStats 注册统计HTTP处理器
func (s *GRPCServerStats) RegisterStats(handler http.Handler) {
	prometheus.MustRegister(s.requestCount, s.requestDuration, s.keepAliveCounter)
	handler.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		promhttp.WriteTo(w, promhttp.Handler().ServeHTTP)
	})
}
