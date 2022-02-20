package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/betalixt/gottp/logger"
)

type HttpClient struct {
	client http.Client
	logger ILogger
}

func (httpClient *HttpClient) Get(
	headers map[string]string,
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.action(
		"GET",
		headers,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) Post(
	headers map[string]string,
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.action(
		"POST",
		headers,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) Patch(
	headers map[string]string,
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.action(
		"PATCH",
		headers,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) Put(
	headers map[string]string,
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.action(
		"PUT",
		headers,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) Delete(
	headers map[string]string,
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.action(
		"DELETE",
		headers,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) PostBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.actionBody(
		"POST",
		headers,
		body,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) PatchBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.actionBody(
		"PATCH",
		headers,
		body,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) PutBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.actionBody(
		"PUT",
		headers,
		body,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) DeleteBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	params ...string,
) (*Response, error) {
	return httpClient.actionBody(
		"DELETE",
		headers,
		body,
		endpoint,
		params...,
	)
}

func (httpClient *HttpClient) action(
	method string,
	headers map[string]string,
	endpoint string,
	pthParms ...string,
) (*Response, error) {
	endpoint, err := formatEp(endpoint, pthParms...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

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

func (httpClient *HttpClient) actionBody(
	method string,
	headers map[string]string,
	body interface{},
	endpoint string,
	pthParms ...string,
) (*Response, error) {
	endpoint, err := formatEp(endpoint, pthParms...)
	if err != nil {
		return nil, err
	}
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

func formatEp(
	format string,
	pthParms ...string,
) (string, error) {
	end := len(format)
	prmCnt := len(pthParms)
	pthNum := 0
	var buffer []byte
	i := 0
	prev := 0
	for i < end {
		for i < end && format[i] != '{' {
			i++
		}
		if i == end {
			break
		}
		if format[i+1] != '}' {
			return "", fmt.Errorf("illegal character/Invalid format in url")
		}
		if pthNum >= prmCnt {
			return "", fmt.Errorf("not enough parameters provided")
		}
		// TODO Maybe can be done in one go?
		escaped := url.QueryEscape(pthParms[pthNum])
		buffer = append(buffer, format[prev:i]...)
		buffer = append(buffer, escaped...)
		pthNum++
		i += 2
		prev = i
	}
	if pthNum != prmCnt {
		return "", fmt.Errorf("too many parameters provided")
	}
	if prev < end {
		buffer = append(buffer, format[prev:end]...)
	}
	return string(buffer), nil
}

// - "Constructors"
func NewClient() *HttpClient {
	client := HttpClient{
		client: *http.DefaultClient,
		logger: logger.NewDefaultLogger(),
	}
	return &client
}
