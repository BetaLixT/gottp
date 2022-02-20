package logger

import "log"

type DefaultLogger struct {
}

func (logger *DefaultLogger) Inf(str string) {
	log.Println(str)
}
func (logger *DefaultLogger) Err(str string) {
	log.Println(str)
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}
