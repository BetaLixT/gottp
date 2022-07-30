# Gottp Http Client
Gottp is a wrapper over the base http's client functionality written in go,
providing additional features like tracing and resilience

Note: This version is built with the github.com/mailru/easyjson package or
similar packages in mind and requires for the request and response bodies to
implement the IRequest and IResponse interfaces respectively, visit the 
versions/base branch to see a version that does not rely on this and uses the
encoding/* packages instead


## Installation
1. Install module
```sh
go get github.com/betalixt/gottp
```
2. Import
```go
import "github.com/betalixt/gottp" 
```
## TODO
* Handle Form Data with files
* Benchmarking and optimizations
* Further Documentation
