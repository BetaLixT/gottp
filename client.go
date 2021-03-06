package gottp

import (
	"bytes"
	"encoding/json"
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

type HttpClient struct {
	client  IInternalClient
	tracer  ITracer
	headers map[string]string
	tid     string
	pid     string
	flg     string
	optn    *ClientOptions
	retr    *retrier.Retrier
}

func (HttpClient *HttpClient) Get(
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		"GET",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

func (HttpClient *HttpClient) Post(
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		"POST",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

func (HttpClient *HttpClient) Patch(
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		"PATCH",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

func (HttpClient *HttpClient) Put(
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		"PUT",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

func (HttpClient *HttpClient) Delete(
	headers map[string]string,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.action(
		"DELETE",
		headers,
		endpoint,
		qParam,
		params...,
	)
}

func (HttpClient *HttpClient) PostBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		"POST",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PatchBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		"PATCH",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PutBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		"PUT",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) DeleteBody(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionBody(
		"DELETE",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PostXml(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		"POST",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PatchXml(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		"PATCH",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PutXml(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		"PUT",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) DeleteXml(
	headers map[string]string,
	body interface{},
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionXML(
		"DELETE",
		headers,
		body,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PostForm(
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
		"POST",
		headers,
		form,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PatchForm(
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
		"PATCH",
		headers,
		form,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) PutForm(
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
		"PUT",
		headers,
		form,
		endpoint,
		qParam, params...,
	)
}

func (HttpClient *HttpClient) DeleteForm(
	headers map[string]string,
	form url.Values,
	endpoint string,
	qParam map[string][]string,
	params ...string,
) (*Response, error) {
	return HttpClient.actionForm(
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
) *HttpClient{
	return &HttpClient{
		client:  client.client,
		tracer:  client.tracer,
		headers: client.headers,
		tid:     client.tid,
		pid:     client.pid,
		flg:     client.flg,
		optn:    optn,
	  retr:    retrier.New(
			retrier.ExponentialBackoff(
				optn.Retry.RetryCount,
				optn.Retry.InitialBackoff,
			),
			retrier.DefaultClassifier{},
		),
	}
}

func (HttpClient *HttpClient) action(
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

	resp, err := HttpClient.runRequest(req)
	if err != nil {
		return nil, err
	}
	respObj := Response(*resp)
	return &respObj, nil
}

func (HttpClient *HttpClient) actionBody(
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

	byts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewReader(byts))
	if err != nil {
		return nil, err
	}

	HttpClient.formHeaders(req, headers)
	req.Header.Set("Content-Type", "application/json")

	resp, err := HttpClient.runRequest(req)
	if err != nil {
		return nil, err
	}
	respObj := Response(*resp)
	return &respObj, nil
}

func (HttpClient *HttpClient) actionXML(
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

	resp, err := HttpClient.runRequest(req)
	if err != nil {
		return nil, err
	}
	respObj := Response(*resp)
	return &respObj, nil
}

func (HttpClient *HttpClient) actionForm(
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

	resp, err := HttpClient.runRequest(req)
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

func (HttpClient *HttpClient) runRequest(
	req *http.Request,
) (*http.Response, error) {
	sid, err := hlpr.GenerateParentId()
	if err == nil {
		req.Header.Add(
			"traceparent",
			fmt.Sprintf("00-%s-%s-%s", HttpClient.tid, sid, HttpClient.flg),
		)
	} else {
		req.Header.Add(
			"traceparent",
			fmt.Sprintf(
				"00-%s-%s-%s",
				HttpClient.tid,
				HttpClient.pid,
				HttpClient.flg,
			),
		)
	}
	start := time.Now()
	var resp *http.Response
	if HttpClient.optn.Retry.Enabled {
		HttpClient.retr.Run(func() error {
			resp, err = HttpClient.client.Do(req)
			if err == nil && resp.StatusCode > 299 {
				for _, val := range HttpClient.optn.Retry.RetriableCodes {
					if resp.StatusCode == val {
						return errors.New("")
					}
				}
			}
			return nil
		})
	} else {
    resp, err = HttpClient.client.Do(req)
	}
	end := time.Now()

	if err != nil {
		HttpClient.tracer.TraceDependency(
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
	HttpClient.tracer.TraceDependency(
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
			"method": req.Method,
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
	tid string,
	pid string,
	flg string,
) *HttpClient {
	optn := DefaultOptions()
	return &HttpClient{
		client:  http.DefaultClient,
		tracer:  tracer,
		headers: headers,
		tid:     tid,
		pid:     pid,
		flg:     flg,
		optn:    optn,
		retr:    retrier.New(
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
		tid:     tid,
		pid:     pid,
		flg:     flg,
		optn:    optn,
		retr:    retrier.New(
			retrier.ExponentialBackoff(
				optn.Retry.RetryCount,
				optn.Retry.InitialBackoff,
			),
			retrier.DefaultClassifier{},
		),
	}
}

//-------
