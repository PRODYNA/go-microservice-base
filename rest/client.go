package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/prodyna/go-microservice-base/trace"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func NewRequest(ctx context.Context, method string, url string) (*http.Request, error) {

	req, err := http.NewRequestWithContext(ctx, method, url, nil)

	if err != nil {
		return nil, err
	}

	// TODO check type
	traceCtx := ctx.Value(trace.TraceContextKey).(trace.TraceContext)

	req.Header.Set(trace.SpanId, traceCtx.SpanId)
	req.Header.Set(trace.ParentSpanId, traceCtx.ParentSpanId)
	req.Header.Set(trace.TraceId, traceCtx.TraceId)
	req.Header.Set(trace.Sampled, traceCtx.Sampled)

	return req, nil
}

type RestClient struct {
	Client       *http.Client
	Method       string
	Url          string
	MinStatus    int
	ParameterMap map[string]string
	HeaderMap    map[string]string
	BodyReader   io.Reader
	Context      context.Context
	Error        error
}

func NewRestClient(client *http.Client, method, url string) *RestClient {
	return &RestClient{
		Client:       client,
		Method:       method,
		Url:          url,
		ParameterMap: make(map[string]string),
		HeaderMap:    make(map[string]string),
		MinStatus:    http.StatusOK,
	}
}

func NewRestClientCtx(client *http.Client, method, url string, ctx context.Context) *RestClient {
	return &RestClient{
		Client:       client,
		Method:       method,
		Url:          url,
		ParameterMap: make(map[string]string),
		HeaderMap:    make(map[string]string),
		MinStatus:    http.StatusOK,
		Context:      ctx,
	}
}

func (r *RestClient) Status(status int) *RestClient {
	r.MinStatus = status
	return r
}

func (r *RestClient) Parameter(key string, value interface{}) *RestClient {
	r.ParameterMap[key] = fmt.Sprintf("%v", value)
	return r
}

func (r *RestClient) Header(key string, value interface{}) *RestClient {
	r.HeaderMap[key] = fmt.Sprintf("%v", value)
	return r
}

func (r *RestClient) BodyString(body string) *RestClient {
	r.BodyReader = strings.NewReader(body)
	return r
}

func (r *RestClient) BodyJson(data interface{}) *RestClient {
	b, err := json.Marshal(data)
	r.HeaderMap["Content-Type"] = "application/json"
	if err == nil {
		r.BodyReader = bytes.NewReader(b)
	}
	return r
}

func (r *RestClient) BodyXML(data interface{}) *RestClient {
	b, err := xml.Marshal(data)
	if err == nil {
		r.BodyReader = bytes.NewReader(b)
	}
	return r
}

func (r *RestClient) Execute() error {
	return r.ExecuteBody(nil)
}

func (r *RestClient) ExecuteBody(result interface{}) error {

	if r.Error != nil {
		return r.Error
	}

	url, err := appendQuery(r.Url, r.ParameterMap)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(r.Method, url, r.BodyReader)

	if err != nil {
		return err
	}

	for k, v := range r.HeaderMap {
		req.Header.Set(k, v)
	}

	res, err := r.Client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode > r.MinStatus {
		return errors.New(fmt.Sprintf("Status is %d, expected is %d", res.StatusCode, r.MinStatus))
	}

	if res.Body == nil {
		return errors.New("Unexpected nil body")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(data) == 0 && res.StatusCode != http.StatusNoContent {
		return errors.New("Unexpected empty body")
	}

	if result != nil {
		err = json.Unmarshal(data, result)
		if err != nil {
			return err
		}
	}

	return nil
}

func appendQuery(u string, parameterMap map[string]string) (string, error) {

	url, err := url.Parse(u)

	if err != nil {
		return "", err
	}

	values := url.Query()

	for k, v := range parameterMap {
		values.Set(k, v)
	}

	url.RawQuery = values.Encode()
	return url.String(), nil

}
