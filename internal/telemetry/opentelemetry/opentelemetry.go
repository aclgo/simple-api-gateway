package opentelemetry

import (
	"context"
	"os"

	"github.com/aclgo/simple-api-gateway/internal/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

type Otel struct {
	provider *sdktrace.TracerProvider
	tracer   trace.Tracer
}

func NewJeager(ctx context.Context, svcName string) (*Otel, error) {
	var (
		tp  *sdktrace.TracerProvider
		err error
	)

	tp, err = create(ctx, svcName)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	tracer := tp.Tracer(svcName)

	return &Otel{
		provider: tp,
		tracer:   tracer,
	}, nil
}

func (o *Otel) Start(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, telemetry.Span) {
	if len(opts) == 0 {
		return o.tracer.Start(ctx, name)
	}

	return o.tracer.Start(ctx, name, opts...)
}

func (o *Otel) Shutdown(ctx context.Context) {
	o.provider.Shutdown(ctx)
}

func create(ctx context.Context, svcName string) (*sdktrace.TracerProvider, error) {
	opts := resource.WithAttributes(
		semconv.ServiceNameKey.String(svcName),
	)

	res, err := resource.New(ctx, opts)
	if err != nil {
		return nil, err
	}

	exp, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(os.Getenv("")),
	)

	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(svcName),
			semconv.DeploymentEnvironmentKey.String("prod"),
		)),
	)

	return tp, nil
}
