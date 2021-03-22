package trace

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateTraceContext(t *testing.T) {
	tx := CreateTraceContext(http.Header{})

	assert.NotNil(t, tx)

}
