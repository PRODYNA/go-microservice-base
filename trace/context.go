package trace

import (
	"context"
	"net/http"
)

const (
	TraceContextKey = "TraceContext"
	TraceId      = "B3-TraceId"
	SpanId       = "B3-SpanId"
	ParentSpanId = "B3-ParentSpanId"
	Sampled      = "B3-Sampled"
)

type TraceContext struct {
	TraceId      string
	SpanId       string
	ParentSpanId string
	Sampled      string
}

func CreateTraceContext(h http.Header) context.Context {

	tx := TraceContext{
		TraceId:      h.Get(TraceId),
		SpanId:       h.Get(SpanId),
		ParentSpanId: h.Get(ParentSpanId),
		Sampled:      h.Get(Sampled),
	}

	ctx := context.WithValue(context.Background(), TraceContextKey, tx)

	return ctx
}
