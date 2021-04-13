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
	Nested   Nested `yaml:"nested"`
}

type Nested struct {
	Pwd string `yaml:"pwd"`
}

var data = `
url: test
userId: userid
password: ${ENV:SECRET_PWD}
nested:
  pwd: ${env:NESTED_SECRET_PWD}
`

func TestLoadConfig(t *testing.T) {

	os.Setenv("SECRET_PWD", "secure")
	os.Setenv("NESTED_SECRET_PWD", "moresecure")

	c := &Configuration{}
	LoadConfig([]byte(data), c)
	assert.Equal(t, "secure", c.Password)
	assert.Equal(t, "moresecure", c.Nested.Pwd)

}


func Test_YamlFile(t *testing.T) {

	b := YamlFile("config.yml")
	assert.NotNil(t, b)
	
}
