package mux

import (
	"github.com/prodyna/go-microservice-base/log"
	"net/http"
	"strings"
)

const (
	KeyPathVariables = "PathVariables"
)

type Mux struct {
	Log   *log.Logger
	Rules map[string][]MuxRule
}

type MuxRule struct {
	Method      string
	ContentType string
	Template    []string
	Handler     http.Handler
}

func NewMux(log *log.Logger) *Mux {
	return &Mux{
		Log:   log,
		Rules: make(map[string][]MuxRule),
	}
}

func (m *Mux) Register(method, path, contentType string, handler http.Handler) {

	key := method + contentType
	var rules []MuxRule

	if _, ok := m.Rules[key]; ok {
		rules = m.Rules[key]
	} else {
		rules = make([]MuxRule, 0)
	}

	mr := MuxRule{
		Method:      method,
		ContentType: contentType,
		Template:    Split(path),
		Handler:     handler,
	}

	rules = append(rules, mr)
	m.Rules[key] = rules

}

func Split(path string) []string {
	s := strings.TrimRight(path, "/")
	return strings.Split(s, "/")
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// set context  r.WithContext()...

	key := r.Method + r.Header.Get("Content-Type")
	path := strings.Split(r.URL.Path, "/")
	ruleList := m.Rules[key]

	if ruleList == nil {
		handleNotFound(w)
		return
	}

	for _, entry := range ruleList {
		match, _ := MatchTemplate(path, entry.Template)
		if match {
			// TODO set vars
			entry.Handler.ServeHTTP(w, r)
			break
		}
	}

	handleNotFound(w)
}

func handleNotFound(w http.ResponseWriter) {
	body := `{ "error" : "Not Found" }`
	w.Write([]byte(body))
	w.WriteHeader(http.StatusNotFound)
}

func PathVariables(r *http.Request) map[string]string {
	m := r.Context().Value(KeyPathVariables).(map[string]string)
	if m != nil {
		return m
	}
	return make(map[string]string)
}
