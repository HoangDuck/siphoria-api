package rotate

import "gopkg.in/natefinch/lumberjack.v2"

const (
	MaxSize    = 1 // MB
	MaxAge     = 7
	MaxBackups = 3
	Compress   = true
	LocalTime  = true
)

type LumberjackSink struct {
	*lumberjack.Logger
}

func (LumberjackSink) Sync() error { return nil }

func Logger(name string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   name,
		MaxSize:    MaxSize,
		MaxAge:     MaxAge,
		MaxBackups: MaxBackups,
		LocalTime:  LocalTime,
		Compress:   Compress,
	}
}
