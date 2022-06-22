package logger

type NoLogger struct {
}

func (logger *NoLogger) Inf(log string) {}
func (logger *NoLogger) Err(log string) {}
