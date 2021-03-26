package error

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	e := NewError("Error", 500)
	value := fmt.Sprintf("%s", e)
	assert.Equal(t, "message: Error, MinStatus 500", value)
}

func TestErrorJson(t *testing.T) {
	e := NewError("Error", 500)
	value := fmt.Sprintf("%s", e.(ServiceError).ErrorJson())
	assert.Equal(t, "{ message: \"Error\", MinStatus: 500 }", value)
}

func TestErrorJsonBytes(t *testing.T) {
	e := NewError("Error", 500)
	value := string(e.(ServiceError).ErrorJsonBytes())
	assert.Equal(t, "{ message: \"Error\", MinStatus: 500 }", value)
}

func TestBadGateway(t *testing.T) {
	e := BadGateway("Error")
	assert.Equal(t, "message: Error, MinStatus 502", e.Error())
}

func TestBadRequest(t *testing.T) {
	e := BadRequest("Error")
	assert.Equal(t, "message: Error, MinStatus 400", e.Error())
}

func TestUnauthorized(t *testing.T) {
	e := Unauthorized("Error")
	assert.Equal(t, "message: Error, MinStatus 401", e.Error())
}

func TestInternalServerError(t *testing.T) {
	e := InternalServerError("Error")
	assert.Equal(t, "message: Error, MinStatus 500", e.Error())
}

func TestNotImplemented(t *testing.T) {
	e := NotImplemented("Error")
	assert.Equal(t, "message: Error, MinStatus 501", e.Error())
}

func TestServiceUnavailable(t *testing.T) {
	e := ServiceUnavailable("Error")
	assert.Equal(t, "message: Error, MinStatus 503", e.Error())
}

func TestForbidden(t *testing.T) {
	e := Forbidden("Error")
	assert.Equal(t, "message: Error, MinStatus 403", e.Error())
}

func TestNotFound(t *testing.T) {
	e := NotFound("Error")
	assert.Equal(t, "message: Error, MinStatus 404", e.Error())
}
