package trace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func StartSpan(ctx context.Context, spanName string) (context.Context, oteltrace.Span) {
	tracer := trace.TracerFromContext(ctx)
	start, span := tracer.Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindClient))
	return start, span
}
func SetSpan(span oteltrace.Span, key string, value interface{}) {
	str, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("[setSpan]err:%v", err)
		return
	}
	span.SetAttributes(attribute.Key(key).String(string(str)))
}

func EndSpan(span oteltrace.Span, err error) {
	defer span.End()

	if err == nil {
		span.SetStatus(codes.Ok, "")
		return
	}
	span.SetStatus(codes.Error, err.Error())
	span.RecordError(err)
}
