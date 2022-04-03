package logger

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logs *zap.Logger
var err error

func init() {
	cfg := zap.NewProductionConfig()
	encoderCfg := zap.NewProductionEncoderConfig()

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.StacktraceKey = ""
	cfg.EncoderConfig = encoderCfg

	logs, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	defer logs.Sync()
}

func Info(message string, d ...interface{}) {
	if d == nil {
		logs.Info(message)
	} else {
		details, _ := json.Marshal(d)
		logs.Info(message, zap.String("details", string(details)))
	}

}

func Error(message string, d ...interface{}) {
	if d == nil {
		logs.Error(message)
	} else {
		details, _ := json.Marshal(d)
		logs.Error(message, zap.String("details", string(details)))
	}
}
