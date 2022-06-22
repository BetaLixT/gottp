package gottp

import (
	"net/url"
	"time"
)

type HttpClient interface {
	Get(headers map[string]string, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	Post(headers map[string]string, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	Put(headers map[string]string, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	Patch(headers map[string]string, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	Delete(headers map[string]string, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	PostBody(headers map[string]string, body interface{}, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	PatchBody(headers map[string]string, body interface{}, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	PutBody(headers map[string]string, body interface{}, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	DeleteBody(headers map[string]string, body interface{}, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	PostForm(headers map[string]string, form url.Values, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	PatchForm(headers map[string]string, form url.Values, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	PutForm(headers map[string]string, form url.Values, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
	DeleteForm(headers map[string]string, form url.Values, endpoint string, qParam map[string][]string, params ...string) (*Response, error)
}

type ITracer interface {
	TraceRequest(isParent bool, method string, path string, query string, statusCode int, bodySize int, ip string, userAgent string, startTimestamp time.Time, eventTimestamp time.Time, fields map[string]string)
	TraceDependency(spanId string, dependencyType string, serviceName string, commandName string, success bool, startTimestamp time.Time, eventTimestamp time.Time, fields map[string]string)
}

type AppInsightsTelemetry interface {
	Close()
}
