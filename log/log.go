package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger initializes zap log
func InitLogger(logLevel, mode string) error {

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

	var config zap.Config
	if mode == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.Level = zap.NewAtomicLevelAt(level)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	Logger, err = config.Build()
	if err != nil {
		return err
	}

	return nil
}
