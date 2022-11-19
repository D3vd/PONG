package log

import (
	"pong/config"

	"go.uber.org/zap"
)

var L *zap.Logger

func Init() (err error) {
	var logger *zap.Logger

	if config.Env == "prod" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

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

func Error(msg string) {
	L.Error(msg)
}

func Errorf(msg string, fields LogFields) {
	L.Error(msg, fields.GetInlineObj())
}

func Request(fields RequestFields) {
	L.Info("", fields.GetInlineObj())
}
