package level

import "go.uber.org/zap/zapcore"

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
	Fatal Level = "fatal"
	Panic Level = "panic"
)

func ZapLevel(lvl Level) zapcore.Level {
	switch lvl {
	case Panic:
		return zapcore.PanicLevel
	case Fatal:
		return zapcore.FatalLevel
	case Error:
		return zapcore.ErrorLevel
	case Warn:
		return zapcore.WarnLevel
	case Info:
		return zapcore.InfoLevel
	case Debug:
		return zapcore.DebugLevel
	default:
		return zapcore.DebugLevel
	}
}
