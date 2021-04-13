package mux

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch(t *testing.T) {

	assert.True(t, Match("/", "/"))
	assert.True(t, Match("/account/1", "account/{id}"))
	assert.True(t, Match("/account/1", "/account/{id}/"))
	assert.True(t, Match("/user/143/account/333/address/333", "/user/{uid}/account/{aid}/address/{adid}"))

	assert.False(t, Match("/user/1", "/user/2"))
	assert.False(t, Match("/user/1", "/user/1/account"))
	assert.False(t, Match("/user/1/account/3", "/user/{uid}/account/{aid}/address"))
}

func Test_MatchTemplatesOk(t *testing.T) {
	path := []string{"api", "v1", "users", "2344"}
	template := []string{"api", "{version}", "users", "{userid}"}

	assert.True(t, MatchTemplate(path, template))
}

func Test_MatchTemplatesNotOk(t *testing.T) {
	path := []string{"api", "v1", "users", "2344"}
	template := []string{"api", "{version}", "list", "{userid}"}

	assert.False(t, MatchTemplate(path, template))
}
