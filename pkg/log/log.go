package log

import (
	"go.uber.org/zap"
)

type Field = zap.Field

var logger *zap.Logger
var module string

func Init() {
	var err error
	const callerSkip = 1
	var options []zap.Option
	options = append(options, zap.AddCallerSkip(callerSkip))
	options = append(options, zap.AddCaller())
	if logger, err = zap.NewProduction(options...); err != nil {
		panic("load logger failed. " + err.Error())
	}
	module = "local-life"
}

func String(key, val string) Field {
	return zap.String(key, val)
}

func Int(key string, val int) Field {
	return zap.Int(key, val)
}

func ErrorF(err error) Field {
	return zap.Error(err)
}

func Debug(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	fields = append([]Field{String("module", module)}, fields...)
	logger.Fatal(msg, fields...)
}
