package jeager

import (
	"context"

	"github.com/aclgo/simple-api-gateway/internal/telemetry"
	"go.opentelemetry.io/otel/trace"
)

type Jeager struct {
}

func NewJeager() *Jeager {
	return &Jeager{}
}

func (j *Jeager) Start(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, telemetry.Span) {
	return nil, nil
}

func (j *Jeager) Shutdown(ctx context.Context) {

}
