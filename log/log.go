package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger initializes zap log
func InitLogger(logLevel string) error {

	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "error":
		level = zap.ErrorLevel
	case "info":
		level = zap.InfoLevel
	default:
		level = zap.InfoLevel
	}

	productionConfig := zap.NewProductionConfig()
	productionConfig.Level = zap.NewAtomicLevelAt(level)
	productionConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	Logger, err = productionConfig.Build()
	if err != nil {
		return err
	}

	return nil
}
