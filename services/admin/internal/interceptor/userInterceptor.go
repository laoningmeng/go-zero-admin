package interceptor

import (
	"context"
	"encoding/json"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	span := trace.SpanFromContext(ctx)
	defer func() {
		if err := recover(); err != nil {
			// todo: alarm program
		}
	}()
	r, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	span.SetAttributes(attribute.KeyValue{
		Key:   "request",
		Value: attribute.StringValue(string(r)),
	})
	resp, err = handler(ctx, req)

	if resp != nil {
		respStr, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		span.SetAttributes(attribute.KeyValue{
			Key:   "response",
			Value: attribute.StringValue(string(respStr)),
		})
	}

	return
}
