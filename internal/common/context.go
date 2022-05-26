package common

import "go.uber.org/zap/zapcore"

const (
	SERVER_CONTEXT_KEY = "server"
)

type ServerContextValue struct {
	Service string
	Version string
	Trace   string
}

func (ctx ServerContextValue) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("service", ctx.Service)
	encoder.AddString("version", ctx.Version)
	encoder.AddString("trace", ctx.Version)
	return nil
}
