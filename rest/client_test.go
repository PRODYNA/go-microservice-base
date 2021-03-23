package rest

import (
	"github.com/prodyna/go-microservice-base/trace"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestNewHttpClient(t *testing.T) {
	h := http.Header{}
	h.Set(trace.Sampled, "1")
	h.Set(trace.TraceId, "36784681")
	h.Set(trace.ParentSpanId, "434242442342423")
	h.Set(trace.SpanId, "1344232")

	ctx := trace.CreateTraceContext(h)

	c := http.Client{Timeout: 10 * time.Second}
	c.Transport = Mock{}

	request,err := NewRequest(ctx, "GET", "https://httpbin.org/headers")
	if err != nil {
		assert.Fail(t, err.Error())
	}

	res,err := c.Do(request)

	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Contains(t, "434242442342423", res.Header.Get(trace.ParentSpanId))
}

type Mock struct {}

func (m Mock) RoundTrip(r *http.Request) (*http.Response, error) {
	res := http.Response{
		Header: r.Header.Clone(),
	}
	return  &res, nil
}


