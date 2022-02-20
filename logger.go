package main

type ILogger interface {
	Inf(string)
	Err(string)
}
