package tracer

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/realHoangHai/kratos/api/gen/go/conf"
	"github.com/realHoangHai/kratos/bootstrap/config"
)

// NewTracerExporter create tracer exporter: zipkin, otlp-http, otlp-grpc
func NewTracerExporter(exporterName, endpoint string, insecure bool) (tracesdk.SpanExporter, error) {
	ctx := context.Background()

	switch exporterName {
	case "zipkin":
		return NewZipkinExporter(ctx, endpoint)
	case "jaeger":
		return nil, errors.New("jaeger exporter is no longer supported, please use otlp-http or otlp-grpc replace it")
	case "otlp-http":
		return NewOtlpHttpExporter(ctx, endpoint, insecure)
	case "otlp-grpc":
		return NewOtlpGrpcExporter(ctx, endpoint, insecure)
	default:
		fallthrough
	case "stdout":
		return stdouttrace.New()
	}
}

// NewTracerProvider create new tracer provider
func NewTracerProvider(cfg *conf.Tracer, serviceInfo *config.ServiceInfo) error {
	if cfg == nil {
		return errors.New("tracer config is nil")
	}

	if !cfg.Enable {
		return nil
	}

	if cfg.Sampler == 0 {
		cfg.Sampler = 1.0
	}

	if cfg.Env == "" {
		cfg.Env = "dev"
	}

	opts := []tracesdk.TracerProviderOption{
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(cfg.GetSampler()))),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceInfo.Name),
			semconv.ServiceVersionKey.String(serviceInfo.Version),
			semconv.ServiceInstanceIDKey.String(serviceInfo.ID),
			attribute.String("env", cfg.GetEnv()),
		)),
	}

	if len(cfg.GetEndpoint()) > 0 {
		exp, err := NewTracerExporter(cfg.GetBatcher(), cfg.GetEndpoint(), cfg.GetInsecure())
		if err != nil {
			panic(err)
		}

		opts = append(opts, tracesdk.WithBatcher(exp))
	}

	tp := tracesdk.NewTracerProvider(opts...)
	if tp == nil {
		return errors.New("create tracer provider failed")
	}

	otel.SetTracerProvider(tp)

	return nil
}

// NewZipkinExporter 创建一个zipkin导出器，默认对端地址：http://localhost:9411/api/v2/spans
func NewZipkinExporter(_ context.Context, endpoint string) (tracesdk.SpanExporter, error) {
	return zipkin.New(endpoint)
}

//// NewJaegerExporter 创建一个jaeger导出器，默认对端地址：http://localhost:14268/api/traces
//func NewJaegerExporter(_ context.Context, endpoint string) (tracesdk.SpanExporter, error) {
//	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
//}

// NewOtlpHttpExporter Create OTLP/HTTP exporter, default port: 4318
func NewOtlpHttpExporter(ctx context.Context, endpoint string, insecure bool, options ...otlptracehttp.Option) (tracesdk.SpanExporter, error) {
	var opts []otlptracehttp.Option
	opts = append(opts, otlptracehttp.WithEndpoint(endpoint))

	if insecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	opts = append(opts, options...)

	return otlptrace.New(
		ctx,
		otlptracehttp.NewClient(opts...),
	)
}

// NewOtlpGrpcExporter Create OTLP/gRPC exporter, default port: 4317
func NewOtlpGrpcExporter(ctx context.Context, endpoint string, insecure bool, options ...otlptracegrpc.Option) (tracesdk.SpanExporter, error) {
	var opts []otlptracegrpc.Option
	opts = append(opts, otlptracegrpc.WithEndpoint(endpoint))

	if insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	opts = append(opts, options...)

	return otlptrace.New(
		ctx,
		otlptracegrpc.NewClient(opts...),
	)
}
