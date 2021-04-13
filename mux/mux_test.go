package mux

import (
	"github.com/prodyna/go-microservice-base/log"
	"net/http"
	"testing"
)

type GetTest struct{}

func (h GetTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func TestNewMux(t *testing.T) {

	m := NewMux(log.RootLogger("test", log.INFO))
	m.Register("GET", "/test/{id}", "", GetTest{})

	r, _ := http.NewRequest("GET", "/test/1334", nil)
	m.ServeHTTP(TestResponseWriter{}, r)

}

type TestResponseWriter struct {
}

func (w TestResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w TestResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w TestResponseWriter) WriteHeader(statusCode int) {}
