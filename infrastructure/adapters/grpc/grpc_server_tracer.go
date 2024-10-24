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

// GRPCServerTracer gRPC服务器追踪器
type GRPCServerTracer struct {
	tracer opentracing.Tracer
}

// NewGRPCServerTracer 创建新的gRPC服务器追踪器实例
func NewGRPCServerTracer(tracer opentracing.Tracer) *GRPCServerTracer {
	return &GRPCServerTracer{
		tracer: tracer,
	}
}

// Wrap 服务包装函数
func (t *GRPCServerTracer) Wrap(server *grpc.Server) {
	// 添加追踪中间件到每个服务方法
	grpc.ServerMethodMiddleware(server, func(ctx context.Context, fullMethod string, handler grpc.UnaryHandler) grpc.UnaryHandler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			span, ctx := t.tracer.StartSpan("grpc_server", opentracing.ChildOf(ctx))
			defer span.Finish()

			span.SetTag("grpc.method", fullMethod)
			span.SetTag("grpc.remote_addr", ctx.Value("remote_addr"))
			span.SetTag("grpc.request_size", req.Size())
			span.SetTag("grpc.response_size", 0)

			ext.SpanKindSet(span, ext.SpanKindServer)
			ext.MessageBusSet(span, "grpc")

			log.Printf("Received %s request", fullMethod)
			resp, err := handler(ctx, req)
			if err != nil {
				span.SetTag("error", err.Error())
			}

			return resp, err
		}
	})
}
