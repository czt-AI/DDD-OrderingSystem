package grpc

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

// GRPCServerFeatures gRPC服务器功能
type GRPCServerFeatures struct {
	requestCount     *prometheus.Counter
	requestDuration  *prometheus.Histogram
	keepAliveCounter *prometheus.Counter
}

// NewGRPCServerFeatures 创建新的gRPC服务器功能实例
func NewGRPCServerFeatures() *GRPCServerFeatures {
	return &GRPCServerFeatures{
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
func (f *GRPCServerFeatures) Wrap(server *grpc.Server) {
	// 添加功能中间件到每个服务方法
	grpc.ServerMethodMiddleware(server, func(ctx context.Context, fullMethod string, handler grpc.UnaryHandler) grpc.UnaryHandler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			f.requestCount.Inc()
			start := time.Now()

			resp, err := handler(ctx, req)
			duration := time.Since(start)

			f.requestDuration.WithLabelValues(fullMethod).Observe(duration.Seconds())

			log.Printf("Handled %s request in %s, Error: %v", fullMethod, duration, err)
			return resp, err
		}
	})

	// 监听keepalive事件
	go func() {
		for {
			select {
			case <-server.Context().Done():
				return
			case <-time.After(1 * time.Second):
				f.keepAliveCounter.Inc()
			}
		}
	}()
}

// RegisterFeatures 注册功能HTTP处理器
func (f *GRPCServerFeatures) RegisterFeatures(handler http.Handler) {
	prometheus.MustRegister(f.requestCount, f.requestDuration, f.keepAliveCounter)
	handler.HandleFunc("/features", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		promhttp.WriteTo(w, promhttp.Handler().ServeHTTP)
	})
}
