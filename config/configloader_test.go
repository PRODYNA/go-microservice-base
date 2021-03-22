package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Configuration struct {
	Url      string `yaml:"url"`
	UserId   string `yaml:"userId"`
	Password string `yaml:"password"`
}

var data = `
url: test
userId: userid
password: ${ENV:SECRET_PWD}
`

func TestLoadConfig(t *testing.T) {

	os.Setenv("SECRET_PWD", "geheim")

	c := &Configuration{}
	LoadConfig([]byte(data), c)
	assert.Equal(t, "geheim", c.Password)

}
