package main

import (
	"fmt"
	"net/http"

	"github.com/betalixt/gottp/logger"
)

type HttpClient struct {
	client http.Client
	logger ILogger
}

func (httpClient *HttpClient) Get(
	endpoint string,
	headers map[string]string,
) (*Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	httpClient.logger.Inf(fmt.Sprintf("making http request GET %s", endpoint))
	resp, err := httpClient.client.Do(req)
	if err != nil {
		httpClient.logger.Err("failed to make request")
		return nil, err
	}
	httpClient.logger.Inf(fmt.Sprintf("resource responded with statusCode %d", resp.StatusCode))
	respObj := Response(*resp)
	return &respObj, nil
}

// - "Constructors"
func NewClient() *HttpClient {
	client := HttpClient{
		client: *http.DefaultClient,
		logger: logger.NewDefaultLogger(),
	}
	return &client
}
