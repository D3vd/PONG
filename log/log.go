package log

import (
	"go.uber.org/zap"
)

var L *zap.Logger

func Init() error {

	logger, err := zap.NewDevelopment()

	if err != nil {
		return err
	}

	L = logger

	return nil
}

func Info(msg string) {
	L.Info(msg)
}

func Infof(msg string, fields LogFields) {
	L.Info(msg, fields.GetInlineObj())
}

func Errorf(msg string, fields LogFields) {
	L.Error(msg, fields.GetInlineObj())
}
