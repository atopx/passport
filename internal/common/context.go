package common

import "go.uber.org/zap/zapcore"

type key string

const (
	ServerContextKey key = "server"
)

type ServerContextValue struct {
	Service string
	Version string
	Trace   string
}

func (ctx ServerContextValue) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("service", ctx.Service)
	encoder.AddString("version", ctx.Version)
	encoder.AddString("trace", ctx.Trace)
	return nil
}
