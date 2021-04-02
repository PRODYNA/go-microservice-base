package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func LoadConfig(data []byte, cfg interface{}) {

	err := yaml.Unmarshal(data, cfg)

	if err != nil {
		fmt.Print(err)
	}

	resolvePasswords(cfg)
}

func YamlFile(file string) []byte {
	data, err :=  ioutil.ReadFile(file)
	if err != nil {
		return []byte{}
	}
	return data
}

func resolvePasswords(cfg interface{}) {

	v := reflect.ValueOf(cfg)

	resolveEnvironment(v)
}

func resolveEnvironment(v reflect.Value) {

	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		if value.Kind() == reflect.String {
			cv := value.String()
			if strings.HasPrefix(cv, "${ENV:") && strings.HasSuffix(cv, "}") {
				env := extractVariable(cv)
				value.SetString(os.Getenv(env))
			}

		} else if value.Kind() == reflect.Struct {
			resolveEnvironment(value)
		}
	}
}

func extractVariable(s string) string {
	s = strings.Replace(s, "${ENV:", "", 1)
	return strings.Replace(s, "}", "", 1)
}
