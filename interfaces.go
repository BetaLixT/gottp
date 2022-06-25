package gottp

import (
	"time"
	"net/http"
)

type IInternalClient interface {
	Do (*http.Request) (*http.Response, error)
}

type ITracer interface {
	TraceRequest(
		isParent bool,
		method string,
		path string,
		query string,
		statusCode int,
		bodySize int,
		ip string,
		userAgent string,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields map[string]string,
	)
	TraceDependency(
		spanId string,
		dependencyType string,
		serviceName string,
		commandName string,
		success bool,
		startTimestamp time.Time,
		eventTimestamp time.Time,
		fields map[string]string,
	)
}
