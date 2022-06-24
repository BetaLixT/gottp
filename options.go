package gottp

type ClientOptions struct {
  Retry RetryPolicy
}

type RetryPolicy struct {
	Enabled        bool
	RetriableCodes []int
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
    },
  }
}
