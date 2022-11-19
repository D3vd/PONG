package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogFields struct {
	Error error
}

func (l LogFields) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("error", l.Error.Error())
	return nil
}

func (l LogFields) GetInlineObj() zapcore.Field {
	return zap.Inline(l)
}

type RequestFields struct {
	Method     string
	RequestURL string
	IP         string
}

func (r RequestFields) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("method", r.Method)
	enc.AddString("request_url", r.RequestURL)
	enc.AddString("ip", r.IP)

	return nil
}

func (r RequestFields) GetInlineObj() zapcore.Field {
	return zap.Inline(r)
}
