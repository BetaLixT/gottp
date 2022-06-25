package gottp

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type responder func() (*http.Response, error)
type MockClient struct {
  resp responder
}

func (client *MockClient) Do(_ *http.Request) (*http.Response, error){
  return client.resp()
}

type MockTrace struct {

}
func (m *MockTrace)TraceRequest(
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
) {

}
func (m *MockTrace) TraceDependency(
	spanId string,
	dependencyType string,
	serviceName string,
	commandName string,
	success bool,
	startTimestamp time.Time,
	eventTimestamp time.Time,
	fields map[string]string,
){

}


func TestGetRetry(t *testing.T) {
  tries := 0
  failThrice := func() (*http.Response, error) {
    if tries < 4 {
      tries++
      return &http.Response{
        StatusCode: 500,
      }, nil
    }
    return &http.Response{
        StatusCode: 200,
      }, nil

  }

  client := NewHttpClientWithClientProvider(
    &MockClient{
      resp: failThrice,
    },
    &MockTrace{},
    nil,
    "",
    "",
    "",
  )
  res, err := client.Get(
    nil,
    "",
    nil,
  )
  if err != nil {
    fmt.Printf("error encountered making requets: %v", err)
    t.FailNow()
  }
  if res.StatusCode != 200 {
    fmt.Printf("Invalid status code")
    t.FailNow()
  }
}

func TestGetRetryFail(t *testing.T) {
  tries := 0
  failThrice := func() (*http.Response, error) {
    if tries < 6 {
      tries++
      return &http.Response{
        StatusCode: 500,
      }, nil
    }
    return &http.Response{
        StatusCode: 200,
      }, nil

  }

  client := NewHttpClientWithClientProvider(
    &MockClient{
      resp: failThrice,
    },
    &MockTrace{},
    nil,
    "",
    "",
    "",
  )
  res, err := client.Get(
    nil,
    "",
    nil,
  )
  if err != nil {
    fmt.Printf("error encountered making requets: %v", err)
    t.FailNow()
  }
  if res.StatusCode == 200 {
    fmt.Printf("Invalid status code")
    t.FailNow()
  }
}
