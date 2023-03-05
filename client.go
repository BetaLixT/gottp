// Package gottp http client
package gottp

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/BetaLixT/go-resiliency/retrier"
	hlpr "github.com/BetaLixT/gottp/helpers"
)

// HttpClient encapsulating REST functionality
type HttpClient struct {
	client  IInternalClient
	tracer  ITracer
	headers map[string]string
	optn    *ClientOptions
	retr    *retrier.Retrier
}

// Get HTTP GET method
func (HttpClient *HttpClient) Get(
	ctx context.Context,
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		ctx,
		"GET",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

// Post HTTP POST method
func (HttpClient *HttpClient) Post(
	ctx context.Context,
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		ctx,
		"POST",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

// Patch HTTP PATH method
func (HttpClient *HttpClient) Patch(
	ctx context.Context,
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		ctx,
		"PATCH",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

// Put HTTP PUT method
func (HttpClient *HttpClient) Put(
	ctx context.Context,
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		ctx,
		"PUT",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

// Delete HTTP DELETE method
func (HttpClient *HttpClient) Delete(
	ctx context.Context,
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		ctx,
		"DELETE",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

// PostBody HTTP POST method with JSON Body
func (HttpClient *HttpClient) PostBody(
	ctx context.Context,
	headers map[string]string,
	body IJsonDTO,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		ctx,
		"POST",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// PatchBody HTTP PATCH method with JSON Body
func (HttpClient *HttpClient) PatchBody(
	ctx context.Context,
	headers map[string]string,
	body IJsonDTO,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		ctx,
		"PATCH",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// PutBody HTTP PUT method with JSON Body
func (HttpClient *HttpClient) PutBody(
	ctx context.Context,
	headers map[string]string,
	body IJsonDTO,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		ctx,
		"PUT",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// DeleteBody HTTP DELETE method with JSON Body
func (HttpClient *HttpClient) DeleteBody(
	ctx context.Context,
	headers map[string]string,
	body IJsonDTO,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		ctx,
		"DELETE",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// PostXml HTTP POST method with XML Body
func (HttpClient *HttpClient) PostXml(
	ctx context.Context,
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		ctx,
		"POST",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// PatchXml HTTP PATCH method with XML Body
func (HttpClient *HttpClient) PatchXml(
	ctx context.Context,
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		ctx,
		"PATCH",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// PutXml HTTP PUT method with XML Body
func (HttpClient *HttpClient) PutXml(
	ctx context.Context,
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		ctx,
		"PUT",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// DeleteXml HTTP DELETE method with XML Body
func (HttpClient *HttpClient) DeleteXml(
	ctx context.Context,
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		ctx,
		"DELETE",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

// PostForm HTTP POST method with Form Body
func (HttpClient *HttpClient) PostForm(
	ctx context.Context,
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
		ctx,
		"POST",
		headers,
		form,
		endpoint,
		qParam, params...,
	)
}

// PatchForm HTTP PATCH method with Form Body
func (HttpClient *HttpClient) PatchForm(
	ctx context.Context,
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
		ctx,
		"PATCH",
		headers,
		form,
		endpoint,
		qParam, params...,
	)
}

// PutForm HTTP PUT method with Form Body
func (HttpClient *HttpClient) PutForm(
	ctx context.Context,
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
		ctx,
		"PUT",
		headers,
		form,
		endpoint,
		qParam, params...,
	)
}

// DeleteForm HTTP DELETE method with Form Body
func (HttpClient *HttpClient) DeleteForm(
	ctx context.Context,
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
		ctx,
		"DELETE",
		headers,
		form,
		endpoint,
		qParam,
		params...,
	)
}

func (client *HttpClient) WithOptions(
	optn *ClientOptions,
) *HttpClient {
	return &HttpClient{
		client:  client.client,
		tracer:  client.tracer,
		headers: client.headers,
		optn:    optn,
		retr: retrier.New(
			retrier.ExponentialBackoff(
				optn.Retry.RetryCount,
				optn.Retry.InitialBackoff,
			),
			retrier.DefaultClassifier{},
		),
	}
}

func (HttpClient *HttpClient) action(
	ctx context.Context,
	method string,
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	pthParms ...string,
) (*Response, error) {
	endpoint, err := formatEp(endpoint, qParam, pthParms...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	HttpClient.formHeaders(req, headers)

	resp, err := HttpClient.runRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	respObj := Response(*resp)
	return &respObj, nil
}

func (client *HttpClient) actionBody(
	ctx context.Context,
	method string,
	headers map[string]string,
	body IJsonDTO,
	endpoint string,
	qParam map[string][]string,
	pthParms ...string,
) (*Response, error) {
	endpoint, err := formatEp(endpoint, qParam, pthParms...)
	if err != nil {
		return nil, err
	}

	byts, err := body.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewReader(byts))
	if err != nil {
		return nil, err
	}

	client.formHeaders(req, headers)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.runRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	respObj := Response(*resp)
	return &respObj, nil
}

func (HttpClient *HttpClient) actionXML(
	ctx context.Context,
	method string,
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	pthParms ...string,
) (*Response, error) {
	endpoint, err := formatEp(endpoint, qParam, pthParms...)
	if err != nil {
		return nil, err
	}

	byts, err := xml.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewReader(byts))
	if err != nil {
		return nil, err
	}

	HttpClient.formHeaders(req, headers)
	req.Header.Set("Content-Type", "application/xml")

	resp, err := HttpClient.runRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	respObj := Response(*resp)
	return &respObj, nil
}

func (HttpClient *HttpClient) actionForm(
	ctx context.Context,
	method string,
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	pthParms ...string,
) (*Response, error) {
	endpoint, err := formatEp(endpoint, qParam, pthParms...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		method,
		endpoint,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return nil, err
	}

	HttpClient.formHeaders(req, headers)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := HttpClient.runRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	respObj := Response(*resp)
	return &respObj, nil
}

func (HttpClient *HttpClient) formHeaders(
	req *http.Request,
	headers map[string]string,
) {
	for key, value := range HttpClient.headers {
		req.Header.Add(key, value)
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}

func (client *HttpClient) runRequest(
	ctx context.Context,
	req *http.Request,
) (*http.Response, error) {
	sid, err := hlpr.GenerateParentId()
	ver, tid, _, rid, flg := client.tracer.ExtractTraceInfo(ctx)
	if err == nil {
		req.Header.Add(
			"traceparent",
			fmt.Sprintf("%s-%s-%s-%s", ver, tid, sid, flg),
		)
	} else {
		req.Header.Add(
			"traceparent",
			fmt.Sprintf(
				"%s-%s-%s-%s",
				ver,
				tid,
				rid,
				flg,
			),
		)
	}
	start := time.Now()
	var resp *http.Response
	if client.optn.Retry.Enabled {
		client.retr.Run(func() error {
			resp, err = client.client.Do(req)
			if err == nil && resp.StatusCode > 299 {
				for _, val := range client.optn.Retry.RetriableCodes {
					if resp.StatusCode == val {
						return errors.New("")
					}
				}
			}
			return nil
		})
	} else {
		resp, err = client.client.Do(req)
	}
	end := time.Now()

	if err != nil {
		client.tracer.TraceDependency(
			ctx,
			sid,
			"http",
			req.URL.Hostname(),
			fmt.Sprintf("%s %s", req.Method, req.URL.RequestURI()),
			false,
			start,
			end,
			// types.NewField("method", req.Method),
			// types.NewField("error", err.Error()),
			map[string]string{"method": req.Method, "error": err.Error()},
		)
		return nil, err
	}
	client.tracer.TraceDependency(
		ctx,
		sid,
		"http",
		req.URL.Hostname(),
		fmt.Sprintf("%s %s", req.Method, req.URL.RequestURI()),
		resp.StatusCode > 199 && resp.StatusCode < 300,
		start,
		end,
		// types.NewField("method", req.Method),
		// types.NewField("statusCode", strconv.Itoa(resp.StatusCode)),
		map[string]string{
			"method":     req.Method,
			"statusCode": strconv.Itoa(resp.StatusCode),
		},
	)
	return resp, err
}

// TODO Pre calculating length and allocating might improve performance
func formatEp(
	format string,
	qParam url.Values,
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

	// TODO found out that url.Values has an Encode funtion that does this,
	//      need to test
	qryBuf := []byte("?")

	for key, vals := range qParam {
		esKey := url.QueryEscape(key)
		for _, val := range vals {
			// TODO just a spike, need to experiment
			query := esKey + "=" + url.QueryEscape(val) + "&"
			qryBuf = append(qryBuf, query...)
		}
	}
	buffer = append(buffer, qryBuf...)
	return string(buffer), nil
}

// - "Constructors"
func NewHttpClientProvider(
	tracer ITracer,
	headers map[string]string,
) *HttpClient {
	optn := DefaultOptions()
	return &HttpClient{
		client:  http.DefaultClient,
		tracer:  tracer,
		headers: headers,
		optn:    optn,
		retr: retrier.New(
			retrier.ExponentialBackoff(
				optn.Retry.RetryCount,
				optn.Retry.InitialBackoff,
			),
			retrier.DefaultClassifier{},
		),
	}
}

func NewHttpClientWithClientProvider(
	client IInternalClient,
	tracer ITracer,
	headers map[string]string,
	tid string,
	pid string,
	flg string,
) *HttpClient {
	optn := DefaultOptions()
	return &HttpClient{
		client:  client,
		tracer:  tracer,
		headers: headers,
		optn:    optn,
		retr: retrier.New(
			retrier.ExponentialBackoff(
				optn.Retry.RetryCount,
				optn.Retry.InitialBackoff,
			),
			retrier.DefaultClassifier{},
		),
	}
}

//-------
