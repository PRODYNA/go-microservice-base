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
	data, err :=  ioutil.ReadFile(getFileName(file))
	if err != nil {
		return []byte{}
	}
	return data
}


func getFileName(file string) string {

	f, err := os.Stat(file)
	if err == nil && !f.IsDir() {
		return file
	}

	fn := "config/" + file
	f, err = os.Stat(fn)
	if err == nil && !f.IsDir() {
		return fn
	}

	fn = "./config/" + file
	f, err = os.Stat(fn)
	if err == nil && !f.IsDir() {
		return fn
	}

	fn = "/config/" + file
	f, err = os.Stat(fn)
	if err == nil && !f.IsDir() {
		return fn
	}

	return file
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
			if strings.HasPrefix(strings.ToLower(cv), "${env:") && strings.HasSuffix(cv, "}") {
				env := extractVariable(cv)
				value.SetString(os.Getenv(env))
			}

		} else if value.Kind() == reflect.Struct {
			resolveEnvironment(value)
		}
	}
}

func extractVariable(s string) string {
	// TODO use regex for replace
	s = strings.Replace(s, "${ENV:", "", 1)
	s = strings.Replace(s, "${Env:", "", 1)
	s = strings.Replace(s, "${env:", "", 1)
	return strings.Replace(s, "}", "", 1)
}
