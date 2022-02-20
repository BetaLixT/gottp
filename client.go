package main

import (
	"encoding/json"
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

func (httpClient *HttpClient) PostBody(
	endpoint string,
	headers map[string]string,
	body interface{},
) (*Response, error) {
	return httpClient.actionBody(
		"POST",
		endpoint,
		headers,
		body,
	)
}

func (httpClient *HttpClient) PatchBody(
	endpoint string,
	headers map[string]string,
	body interface{},
) (*Response, error) {
	return httpClient.actionBody(
		"PATCH",
		endpoint,
		headers,
		body,
	)
}

func (httpClient *HttpClient) PutBody(
	endpoint string,
	headers map[string]string,
	body interface{},
) (*Response, error) {
	return httpClient.actionBody(
		"PUT",
		endpoint,
		headers,
		body,
	)
}

func (httpClient *HttpClient) DeleteBody(
	endpoint string,
	headers map[string]string,
	body interface{},
) (*Response, error) {
	return httpClient.actionBody(
		"DELETE",
		endpoint,
		headers,
		body,
	)
}

func (httpClient *HttpClient) actionBody(
	method string,
	endpoint string,
	headers map[string]string,
	body interface{},
) (*Response, error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	byts, err := json.Marshal(body)
	if err != nil {
		httpClient.logger.Err("failed to marshal body")
		return nil, err
	}
	req.Body.Read(byts)

	httpClient.logger.Inf(fmt.Sprintf("making http request %s %s", method, endpoint))
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
