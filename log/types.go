package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogFields struct {
	Error error `json:"error"`
}

func (l LogFields) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("error", l.Error.Error())
	return nil
}

func (l LogFields) GetInlineObj() zapcore.Field {
	return zap.Inline(l)
}
