package gottp

type ILogger interface {
	Inf(string)
	Err(string)
}
