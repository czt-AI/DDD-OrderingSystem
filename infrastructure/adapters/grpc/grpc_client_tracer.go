package grpc

import (
	"context"
	"log"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc"
)

// GRPCClientTracer gRPC客户端追踪器
type GRPCClientTracer struct {
	tracer opentracing.Tracer
}

// NewGRPCClientTracer 创建新的gRPC客户端追踪器实例
func NewGRPCClientTracer(tracer opentracing.Tracer) *GRPCClientTracer {
	return &GRPCClientTracer{
		tracer: tracer,
	}
}

// WrapClientDialer 包装客户端 dialer，用于添加追踪
func (t *GRPCClientTracer) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
	return grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		span, ctx := t.tracer.StartSpan("grpc_client", opentracing.ChildOf(ctx))
		defer span.Finish()

		span.SetTag("grpc.method", method)
		span.SetTag("grpc.remote_addr", cc.GetAddr())
		span.SetTag("grpc.request_size", req.Size())
		span.SetTag("grpc.response_size", 0)

		ext.SpanKindSet(span, ext.SpanKindClient)
		ext.MessageBusSet(span, "grpc")

		start := time.Now()
		log.Printf("Sending %s request to %s", method, cc.GetAddr())
		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(start)

		if err != nil {
			span.SetTag("error", err.Error())
		}

		log.Printf("Received %s response from %s in %s", method, cc.GetAddr(), duration)
		return err
	})(dialer)
}
