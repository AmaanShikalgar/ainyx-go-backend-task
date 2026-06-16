package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init(env string) {
	var err error

	if env == "production" {

		Log, err = zap.NewProduction()
	} else {

		Log, err = zap.NewDevelopment()
	}

	if err != nil {

		panic("failed to initialize logger: " + err.Error())
	}

	zap.ReplaceGlobals(Log)
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return Log.With(fields...)
}

func Err(err error) zap.Field {
	return zap.Error(err)
}

func Int32(key string, val int32) zap.Field {
	return zap.Int32(key, val)
}

func Str(key string, val string) zap.Field {
	return zap.String(key, val)
}

var _ = zapcore.DebugLevel
