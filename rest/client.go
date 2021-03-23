package rest

import (
	"context"
	"github.com/prodyna/go-microservice-base/trace"
	"net/http"
)

func NewRequest(ctx context.Context, method string, url string) (*http.Request, error) {

	req,err := http.NewRequestWithContext(ctx, method, url, nil)

	if err != nil {
		return nil, err
	}

	// TODO check type
	traceCtx := ctx.Value(trace.TraceContextKey).(trace.TraceContext)

	req.Header.Set(trace.SpanId, traceCtx.SpanId)
	req.Header.Set(trace.ParentSpanId, traceCtx.ParentSpanId)
	req.Header.Set(trace.TraceId, traceCtx.TraceId)
	req.Header.Set(trace.Sampled, traceCtx.Sampled)

	return req, nil
}
