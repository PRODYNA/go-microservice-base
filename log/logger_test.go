package log

import (
	"github.com/prodyna/go-microservice-base/trace"
	"net/http"
	"testing"
)

func TestLogger(t *testing.T) {


	log := RootLoggerCtx("service", ERROR)

	//log.Error("op", "msg")

	l2 := log.Logger(log, "adapter")
	l3 := log.Logger(l2, "xyz")

	//l2.Error("op", "msg", context.Background())


	m := make(map[string]string)
	m["a"] = "b"
	m["b"] = "a"

	h := http.Header{}
	h.Add("X-B3-TraceId", "1234")

	ctx := trace.CreateTraceContext(h)

	l3.Error("op", "msg", m, ctx)
}
