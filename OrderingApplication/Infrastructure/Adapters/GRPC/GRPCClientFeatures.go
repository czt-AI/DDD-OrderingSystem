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

// GRPCClientFeatures gRPC客户端功能
type GRPCClientFeatures struct {
	requestCount     *prometheus.Counter
	requestDuration  *prometheus.Histogram
	keepAliveCounter *prometheus.Counter
}

// NewGRPCClientFeatures 创建新的gRPC客户端功能实例
func NewGRPCClientFeatures() *GRPCClientFeatures {
	return &GRPCClientFeatures{
		requestCount:     prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_client_request_count",
			Help: "Total number of gRPC client requests.",
		}),
		requestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "grpc_client_request_duration_seconds",
			Help:    "Duration of gRPC client requests in seconds.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		}),
		keepAliveCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "grpc_client_keepalive_count",
			Help: "Total number of keepalive pings sent by the client.",
		}),
	}
}

// WrapClientDialer 包装客户端 dialer，用于添加功能
func (f *GRPCClientFeatures) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		f.requestCount.Inc()
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)

		f.requestDuration.WithLabelValues(method).Observe(duration.Seconds())

		log.Printf("Sent %s request to %s in %s, Error: %v", method, cc.GetAddr(), duration, err)
		return err
	})(dialer)
}

// RegisterFeatures 注册功能HTTP处理器
func (f *GRPCClientFeatures) RegisterFeatures(handler http.Handler) {
	prometheus.MustRegister(f.requestCount, f.requestDuration, f.keepAliveCounter)
	handler.HandleFunc("/features", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		promhttp.WriteTo(w, promhttp.Handler().ServeHTTP)
	})
}