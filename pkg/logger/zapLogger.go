package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

type ZapLogger struct {
	log *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	var config zap.Config

	config = zap.NewDevelopmentConfig()
	config.DisableStacktrace = true
	config.EncoderConfig.ConsoleSeparator = "  |  "
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(filepath.Base(caller.FullPath()))
	}
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

	logger, err := config.Build(zap.AddCaller(), zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}

	return &ZapLogger{logger.Sugar()}
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.log.Info(args...)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.log.Error(args...)
}

func (z *ZapLogger) Fatal(args ...interface{}) {
	z.log.Fatal(args...)
}
