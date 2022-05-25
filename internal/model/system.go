package model

import "go.uber.org/zap/zapcore"

type System struct {
	Server  string `json:"server"`
	Version string `json:"version"`
}

func (sys System) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("system", sys.Server)
	encoder.AddString("version", sys.Version)
	return nil
}
