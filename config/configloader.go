package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
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
	return []byte("data")
}

func resolvePasswords(cfg interface{}) {

	v := reflect.ValueOf(cfg)

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

		}
	}
}

func extractVariable(s string) string {
	s = strings.Replace(s, "${ENV:", "", 1)
	return strings.Replace(s, "}", "", 1)
}
