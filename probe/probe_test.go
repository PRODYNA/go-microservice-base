package probe

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLiveProbe(t *testing.T) {
	p := NewProbeService()
	p.AddLive(NewGoRoutineProbe(10))
	assert.True(t, p.verifyLive())
}


func TestReadyProbe(t *testing.T) {
	p := NewProbeService()
	p.AddReady(NewGoRoutineProbe(10))
	assert.True(t, p.verifyReady())
}


func TestStartProbe(t *testing.T) {
	p := NewProbeService()
	p.AddStart(NewGoRoutineProbe(10))
	assert.True(t, p.verifyStart())
}


type MockResponseWriter struct {}


func TestServeHttp(t *testing.T) {
	p := NewProbeService()
	p.AddStart(NewGoRoutineProbe(10))
	p.AddReady(NewGoRoutineProbe(10))
	p.AddLive(NewGoRoutineProbe(10))


	p.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET","http://localhost:8080/ready",strings.NewReader("")))
	p.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET","http://localhost:8080/live",strings.NewReader("")))
	p.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET","http://localhost:8080/start",strings.NewReader("")))


	p.AddLive(NewGoRoutineProbe(0))
	p.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET","http://localhost:8080/live",strings.NewReader("")))

	p.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("GET","http://localhost:8080/XXX",strings.NewReader("")))
}
