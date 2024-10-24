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

// GRPCClientTelemetry gRPC客户端遥测
type GRPCClientTelemetry struct {
	tracer opentracing.Tracer
}

// NewGRPCClientTelemetry 创建新的gRPC客户端遥测实例
func NewGRPCClientTelemetry(tracer opentracing.Tracer) *GRPCClientTelemetry {
	return &GRPCClientTelemetry{
		tracer: tracer,
	}
}

// WrapClientDialer 包装客户端 dialer，用于添加遥测
func (t *GRPCClientTelemetry) WrapClientDialer(dialer grpc.Dialer) grpc.Dialer {
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

		span.SetTag("grpc.duration", duration.Seconds())
		log.Printf("Received %s response from %s in %s", method, cc.GetAddr(), duration)
		return err
	})(dialer)
}