package gottp

import "time"

type ClientOptions struct {
  Retry RetryPolicy
}

type RetryPolicy struct {
	Enabled        bool
	RetriableCodes []int
	RetryCount     int
	InitialBackoff time.Duration
}

func DefaultOptions() *ClientOptions {
  return &ClientOptions{
    Retry: RetryPolicy{
      Enabled: true,
      RetriableCodes: []int {
        408,
        500,
        502,
        503,
        504,
      },
      RetryCount: 5,
      InitialBackoff: 100 * time.Millisecond,
    },
  }
}
